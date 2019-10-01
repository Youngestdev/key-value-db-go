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