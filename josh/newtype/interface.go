package newtype

type Status interface {
	CheckStatus(websiteName string) (status bool, err error)
}
