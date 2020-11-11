// Code generated by "dbusutil-gen -type Network network.go"; DO NOT EDIT.

package network

func (v *Network) setPropVpnEnabled(value bool) (changed bool) {
	if v.VpnEnabled != value {
		v.VpnEnabled = value
		v.emitPropChangedVpnEnabled(value)
		return true
	}
	return false
}

func (v *Network) emitPropChangedVpnEnabled(value bool) error {
	return v.service.EmitPropertyChanged(v, "VpnEnabled", value)
}