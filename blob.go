package git4go

func (r *Repository) LookupBlob(oid *Oid) (*Blob, error) {
	obj, err := objectLookupPrefix(r, oid, GitOidHexSize, ObjectBlob)
	return obj.(*Blob), err
}

func (r *Repository) LookupPrefixBlob(oid *Oid, length int) (*Blob, error) {
	obj, err := objectLookupPrefix(r, oid, length, ObjectBlob)
	return obj.(*Blob), err
}

type Blob struct {
	GitObject
	contents []byte
}

func (b *Blob) Type() ObjectType {
	return ObjectBlob
}

func (b *Blob) Size() int64 {
	return int64(len(b.contents))
}

func (b *Blob) Contents() []byte {
	return b.contents
}

func newBlob(repo *Repository, oid *Oid, contents []byte) *Blob {
	return &Blob{
		contents: contents,
		GitObject: GitObject{
			repo: repo,
			oid:  oid,
		},
	}
}