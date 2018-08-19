package userutil

type InputError struct {
	reason        string
	didNotUseYesNo bool
}

func (o InputError) Error() string {
	return o.reason
}

func (o InputError) YesNoNotProvided() bool {
	return o.didNotUseYesNo
}

type UserError struct {
	reason        string
	notRoot       bool
	unableToCheck bool
}

func (o UserError) Error() string {
	return o.reason
}

func (o UserError) NotRoot() bool {
	return o.notRoot
}

func (o UserError) CheckFailed() bool {
	return o.unableToCheck
}
