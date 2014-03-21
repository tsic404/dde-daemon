/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
        "github.com/howeyc/fsnotify"
        "os"
)

func (op *Manager) listenThemeDir(dir string) {
        //logObject.Infof("Listen Dir: %s\n", dir)
        watcher, err := fsnotify.NewWatcher()
        if err != nil {
                logObject.Infof("Create new watch failed: %v", err)
                return
        }

        if ok, _ := objUtil.IsFileExist(dir); !ok {
                err = os.MkdirAll(dir, 0755)
                if err != nil {
                        logObject.Infof("Make dir '%s' failed: %v", dir, err)
                        return
                }
        }

        err = watcher.Watch(dir)
        if err != nil {
                logObject.Infof("Watch '%s' failed: %v", dir, err)
                return
        }

        go func() {
                defer watcher.Close()
                for {
                        select {
                        case ev := <-watcher.Event:
                                if !ev.IsDelete() {
                                        op.updateAllProps()
                                }
                        case err := <-watcher.Error:
                                logObject.Warningf("Watch Error: %v", err)
                        }
                }
        }()
}

func (op *Manager) listenBackgroundDir(dir string) {
        if ok, _ := objUtil.IsFileExist(dir); !ok {
                err := os.MkdirAll(dir, 0755)
                if err != nil {
                        logObject.Infof("mkdir '%s' failed: %v", dir, err)
                        return
                }
        }

        watcher, err := fsnotify.NewWatcher()
        if err != nil {
                logObject.Infof("Create new watch failed: %v", err)
                return
        }

        go func() {
                defer watcher.Close()
                for {
                        select {
                        case ev := <-watcher.Event:
                                if !ev.IsDelete() {
                                        op.setPropName("BackgroundList")
                                }
                        case err := <-watcher.Error:
                                logObject.Warningf("Watch Error: %v", err)
                        }
                }
        }()
}
