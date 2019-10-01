package KeyValueDB

type ValueObject struct {
	Key, Value string
	hashcode   int
}

func (k ValueObject) KeyValueObject(key, value string) {
	k.Key = key
	k.Value = value
	k.hashcode = 0
}

func (k *ValueObject) Test(x int) {
	k.hashcode = x
}

func (k ValueObject) GetKey() string {
	return k.Key
}

func (k ValueObject) GetVal() string {
	return k.Value
}

func (k *ValueObject) SetKey(key string) {
	k.Key = key
}

func (k *ValueObject) SetVal(val string) {
	k.Value = val
}

func (k *ValueObject) GetH() int {
	return k.hashcode
}

// This isn't functional, I guess.

func (h *ValueObject) HashCode() int {
	if h.hashcode != 0 {
		return h.hashcode
	}

	if len(h.Key) == 0 {
		return h.hashcode
	}
	var x int
	for i := 0; i < len(h.Key); i++ {
		chr := string([]rune(h.Key))
		for _, c := range chr {
			return int(c)
		}
		h.hashcode = h.hashcode<<5 - h.hashcode + x
		h.hashcode |= 0
	}
	return h.hashcode
}

func (k *ValueObject) ToString() string {
	return "KeyValueObject@" + string(k.HashCode()) + ":Key=" + k.Key + ", Value=" + k.Value + ""
}
