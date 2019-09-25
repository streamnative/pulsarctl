package pulsar

type APIVersion int

const (
	V1 APIVersion = iota
	V2
	V3
)

const DefaultAPIVersion = "v2"

func (v APIVersion) String() string {
	switch v {
	case V1:
		return ""
	case V2:
		return "v2"
	case V3:
		return "v3"
	}

	return DefaultAPIVersion
}
