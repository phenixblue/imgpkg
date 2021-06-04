// Code generated by counterfeiter. DO NOT EDIT.
package imagesetfakes

import (
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/k14s/imgpkg/pkg/imgpkg/imageset"
)

type FakeImagesReaderWriter struct {
	DigestStub        func(name.Reference) (v1.Hash, error)
	digestMutex       sync.RWMutex
	digestArgsForCall []struct {
		arg1 name.Reference
	}
	digestReturns struct {
		result1 v1.Hash
		result2 error
	}
	digestReturnsOnCall map[int]struct {
		result1 v1.Hash
		result2 error
	}
	FirstImageExistsStub        func([]string) (string, error)
	firstImageExistsMutex       sync.RWMutex
	firstImageExistsArgsForCall []struct {
		arg1 []string
	}
	firstImageExistsReturns struct {
		result1 string
		result2 error
	}
	firstImageExistsReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetStub        func(name.Reference) (*remote.Descriptor, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 name.Reference
	}
	getReturns struct {
		result1 *remote.Descriptor
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *remote.Descriptor
		result2 error
	}
	ImageStub        func(name.Reference) (v1.Image, error)
	imageMutex       sync.RWMutex
	imageArgsForCall []struct {
		arg1 name.Reference
	}
	imageReturns struct {
		result1 v1.Image
		result2 error
	}
	imageReturnsOnCall map[int]struct {
		result1 v1.Image
		result2 error
	}
	IndexStub        func(name.Reference) (v1.ImageIndex, error)
	indexMutex       sync.RWMutex
	indexArgsForCall []struct {
		arg1 name.Reference
	}
	indexReturns struct {
		result1 v1.ImageIndex
		result2 error
	}
	indexReturnsOnCall map[int]struct {
		result1 v1.ImageIndex
		result2 error
	}
	MultiWriteStub        func(map[name.Reference]remote.Taggable, int, chan v1.Update) error
	multiWriteMutex       sync.RWMutex
	multiWriteArgsForCall []struct {
		arg1 map[name.Reference]remote.Taggable
		arg2 int
		arg3 chan v1.Update
	}
	multiWriteReturns struct {
		result1 error
	}
	multiWriteReturnsOnCall map[int]struct {
		result1 error
	}
	WriteImageStub        func(name.Reference, v1.Image) error
	writeImageMutex       sync.RWMutex
	writeImageArgsForCall []struct {
		arg1 name.Reference
		arg2 v1.Image
	}
	writeImageReturns struct {
		result1 error
	}
	writeImageReturnsOnCall map[int]struct {
		result1 error
	}
	WriteIndexStub        func(name.Reference, v1.ImageIndex) error
	writeIndexMutex       sync.RWMutex
	writeIndexArgsForCall []struct {
		arg1 name.Reference
		arg2 v1.ImageIndex
	}
	writeIndexReturns struct {
		result1 error
	}
	writeIndexReturnsOnCall map[int]struct {
		result1 error
	}
	WriteTagStub        func(name.Tag, remote.Taggable) error
	writeTagMutex       sync.RWMutex
	writeTagArgsForCall []struct {
		arg1 name.Tag
		arg2 remote.Taggable
	}
	writeTagReturns struct {
		result1 error
	}
	writeTagReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImagesReaderWriter) Digest(arg1 name.Reference) (v1.Hash, error) {
	fake.digestMutex.Lock()
	ret, specificReturn := fake.digestReturnsOnCall[len(fake.digestArgsForCall)]
	fake.digestArgsForCall = append(fake.digestArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	stub := fake.DigestStub
	fakeReturns := fake.digestReturns
	fake.recordInvocation("Digest", []interface{}{arg1})
	fake.digestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesReaderWriter) DigestCallCount() int {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	return len(fake.digestArgsForCall)
}

func (fake *FakeImagesReaderWriter) DigestCalls(stub func(name.Reference) (v1.Hash, error)) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = stub
}

func (fake *FakeImagesReaderWriter) DigestArgsForCall(i int) name.Reference {
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	argsForCall := fake.digestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesReaderWriter) DigestReturns(result1 v1.Hash, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	fake.digestReturns = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) DigestReturnsOnCall(i int, result1 v1.Hash, result2 error) {
	fake.digestMutex.Lock()
	defer fake.digestMutex.Unlock()
	fake.DigestStub = nil
	if fake.digestReturnsOnCall == nil {
		fake.digestReturnsOnCall = make(map[int]struct {
			result1 v1.Hash
			result2 error
		})
	}
	fake.digestReturnsOnCall[i] = struct {
		result1 v1.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) FirstImageExists(arg1 []string) (string, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.firstImageExistsMutex.Lock()
	ret, specificReturn := fake.firstImageExistsReturnsOnCall[len(fake.firstImageExistsArgsForCall)]
	fake.firstImageExistsArgsForCall = append(fake.firstImageExistsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.FirstImageExistsStub
	fakeReturns := fake.firstImageExistsReturns
	fake.recordInvocation("FirstImageExists", []interface{}{arg1Copy})
	fake.firstImageExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesReaderWriter) FirstImageExistsCallCount() int {
	fake.firstImageExistsMutex.RLock()
	defer fake.firstImageExistsMutex.RUnlock()
	return len(fake.firstImageExistsArgsForCall)
}

func (fake *FakeImagesReaderWriter) FirstImageExistsCalls(stub func([]string) (string, error)) {
	fake.firstImageExistsMutex.Lock()
	defer fake.firstImageExistsMutex.Unlock()
	fake.FirstImageExistsStub = stub
}

func (fake *FakeImagesReaderWriter) FirstImageExistsArgsForCall(i int) []string {
	fake.firstImageExistsMutex.RLock()
	defer fake.firstImageExistsMutex.RUnlock()
	argsForCall := fake.firstImageExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesReaderWriter) FirstImageExistsReturns(result1 string, result2 error) {
	fake.firstImageExistsMutex.Lock()
	defer fake.firstImageExistsMutex.Unlock()
	fake.FirstImageExistsStub = nil
	fake.firstImageExistsReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) FirstImageExistsReturnsOnCall(i int, result1 string, result2 error) {
	fake.firstImageExistsMutex.Lock()
	defer fake.firstImageExistsMutex.Unlock()
	fake.FirstImageExistsStub = nil
	if fake.firstImageExistsReturnsOnCall == nil {
		fake.firstImageExistsReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.firstImageExistsReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) Get(arg1 name.Reference) (*remote.Descriptor, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesReaderWriter) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeImagesReaderWriter) GetCalls(stub func(name.Reference) (*remote.Descriptor, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeImagesReaderWriter) GetArgsForCall(i int) name.Reference {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesReaderWriter) GetReturns(result1 *remote.Descriptor, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *remote.Descriptor
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) GetReturnsOnCall(i int, result1 *remote.Descriptor, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *remote.Descriptor
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *remote.Descriptor
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) Image(arg1 name.Reference) (v1.Image, error) {
	fake.imageMutex.Lock()
	ret, specificReturn := fake.imageReturnsOnCall[len(fake.imageArgsForCall)]
	fake.imageArgsForCall = append(fake.imageArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	stub := fake.ImageStub
	fakeReturns := fake.imageReturns
	fake.recordInvocation("Image", []interface{}{arg1})
	fake.imageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesReaderWriter) ImageCallCount() int {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return len(fake.imageArgsForCall)
}

func (fake *FakeImagesReaderWriter) ImageCalls(stub func(name.Reference) (v1.Image, error)) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = stub
}

func (fake *FakeImagesReaderWriter) ImageArgsForCall(i int) name.Reference {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	argsForCall := fake.imageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesReaderWriter) ImageReturns(result1 v1.Image, result2 error) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	fake.imageReturns = struct {
		result1 v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) ImageReturnsOnCall(i int, result1 v1.Image, result2 error) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	if fake.imageReturnsOnCall == nil {
		fake.imageReturnsOnCall = make(map[int]struct {
			result1 v1.Image
			result2 error
		})
	}
	fake.imageReturnsOnCall[i] = struct {
		result1 v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) Index(arg1 name.Reference) (v1.ImageIndex, error) {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct {
		arg1 name.Reference
	}{arg1})
	stub := fake.IndexStub
	fakeReturns := fake.indexReturns
	fake.recordInvocation("Index", []interface{}{arg1})
	fake.indexMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesReaderWriter) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeImagesReaderWriter) IndexCalls(stub func(name.Reference) (v1.ImageIndex, error)) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = stub
}

func (fake *FakeImagesReaderWriter) IndexArgsForCall(i int) name.Reference {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	argsForCall := fake.indexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesReaderWriter) IndexReturns(result1 v1.ImageIndex, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 v1.ImageIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) IndexReturnsOnCall(i int, result1 v1.ImageIndex, result2 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 v1.ImageIndex
			result2 error
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 v1.ImageIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesReaderWriter) MultiWrite(arg1 map[name.Reference]remote.Taggable, arg2 int, arg3 chan v1.Update) error {
	fake.multiWriteMutex.Lock()
	ret, specificReturn := fake.multiWriteReturnsOnCall[len(fake.multiWriteArgsForCall)]
	fake.multiWriteArgsForCall = append(fake.multiWriteArgsForCall, struct {
		arg1 map[name.Reference]remote.Taggable
		arg2 int
		arg3 chan v1.Update
	}{arg1, arg2, arg3})
	stub := fake.MultiWriteStub
	fakeReturns := fake.multiWriteReturns
	fake.recordInvocation("MultiWrite", []interface{}{arg1, arg2, arg3})
	fake.multiWriteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImagesReaderWriter) MultiWriteCallCount() int {
	fake.multiWriteMutex.RLock()
	defer fake.multiWriteMutex.RUnlock()
	return len(fake.multiWriteArgsForCall)
}

func (fake *FakeImagesReaderWriter) MultiWriteCalls(stub func(map[name.Reference]remote.Taggable, int, chan v1.Update) error) {
	fake.multiWriteMutex.Lock()
	defer fake.multiWriteMutex.Unlock()
	fake.MultiWriteStub = stub
}

func (fake *FakeImagesReaderWriter) MultiWriteArgsForCall(i int) (map[name.Reference]remote.Taggable, int, chan v1.Update) {
	fake.multiWriteMutex.RLock()
	defer fake.multiWriteMutex.RUnlock()
	argsForCall := fake.multiWriteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImagesReaderWriter) MultiWriteReturns(result1 error) {
	fake.multiWriteMutex.Lock()
	defer fake.multiWriteMutex.Unlock()
	fake.MultiWriteStub = nil
	fake.multiWriteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) MultiWriteReturnsOnCall(i int, result1 error) {
	fake.multiWriteMutex.Lock()
	defer fake.multiWriteMutex.Unlock()
	fake.MultiWriteStub = nil
	if fake.multiWriteReturnsOnCall == nil {
		fake.multiWriteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.multiWriteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteImage(arg1 name.Reference, arg2 v1.Image) error {
	fake.writeImageMutex.Lock()
	ret, specificReturn := fake.writeImageReturnsOnCall[len(fake.writeImageArgsForCall)]
	fake.writeImageArgsForCall = append(fake.writeImageArgsForCall, struct {
		arg1 name.Reference
		arg2 v1.Image
	}{arg1, arg2})
	stub := fake.WriteImageStub
	fakeReturns := fake.writeImageReturns
	fake.recordInvocation("WriteImage", []interface{}{arg1, arg2})
	fake.writeImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImagesReaderWriter) WriteImageCallCount() int {
	fake.writeImageMutex.RLock()
	defer fake.writeImageMutex.RUnlock()
	return len(fake.writeImageArgsForCall)
}

func (fake *FakeImagesReaderWriter) WriteImageCalls(stub func(name.Reference, v1.Image) error) {
	fake.writeImageMutex.Lock()
	defer fake.writeImageMutex.Unlock()
	fake.WriteImageStub = stub
}

func (fake *FakeImagesReaderWriter) WriteImageArgsForCall(i int) (name.Reference, v1.Image) {
	fake.writeImageMutex.RLock()
	defer fake.writeImageMutex.RUnlock()
	argsForCall := fake.writeImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImagesReaderWriter) WriteImageReturns(result1 error) {
	fake.writeImageMutex.Lock()
	defer fake.writeImageMutex.Unlock()
	fake.WriteImageStub = nil
	fake.writeImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteImageReturnsOnCall(i int, result1 error) {
	fake.writeImageMutex.Lock()
	defer fake.writeImageMutex.Unlock()
	fake.WriteImageStub = nil
	if fake.writeImageReturnsOnCall == nil {
		fake.writeImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteIndex(arg1 name.Reference, arg2 v1.ImageIndex) error {
	fake.writeIndexMutex.Lock()
	ret, specificReturn := fake.writeIndexReturnsOnCall[len(fake.writeIndexArgsForCall)]
	fake.writeIndexArgsForCall = append(fake.writeIndexArgsForCall, struct {
		arg1 name.Reference
		arg2 v1.ImageIndex
	}{arg1, arg2})
	stub := fake.WriteIndexStub
	fakeReturns := fake.writeIndexReturns
	fake.recordInvocation("WriteIndex", []interface{}{arg1, arg2})
	fake.writeIndexMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImagesReaderWriter) WriteIndexCallCount() int {
	fake.writeIndexMutex.RLock()
	defer fake.writeIndexMutex.RUnlock()
	return len(fake.writeIndexArgsForCall)
}

func (fake *FakeImagesReaderWriter) WriteIndexCalls(stub func(name.Reference, v1.ImageIndex) error) {
	fake.writeIndexMutex.Lock()
	defer fake.writeIndexMutex.Unlock()
	fake.WriteIndexStub = stub
}

func (fake *FakeImagesReaderWriter) WriteIndexArgsForCall(i int) (name.Reference, v1.ImageIndex) {
	fake.writeIndexMutex.RLock()
	defer fake.writeIndexMutex.RUnlock()
	argsForCall := fake.writeIndexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImagesReaderWriter) WriteIndexReturns(result1 error) {
	fake.writeIndexMutex.Lock()
	defer fake.writeIndexMutex.Unlock()
	fake.WriteIndexStub = nil
	fake.writeIndexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteIndexReturnsOnCall(i int, result1 error) {
	fake.writeIndexMutex.Lock()
	defer fake.writeIndexMutex.Unlock()
	fake.WriteIndexStub = nil
	if fake.writeIndexReturnsOnCall == nil {
		fake.writeIndexReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeIndexReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteTag(arg1 name.Tag, arg2 remote.Taggable) error {
	fake.writeTagMutex.Lock()
	ret, specificReturn := fake.writeTagReturnsOnCall[len(fake.writeTagArgsForCall)]
	fake.writeTagArgsForCall = append(fake.writeTagArgsForCall, struct {
		arg1 name.Tag
		arg2 remote.Taggable
	}{arg1, arg2})
	stub := fake.WriteTagStub
	fakeReturns := fake.writeTagReturns
	fake.recordInvocation("WriteTag", []interface{}{arg1, arg2})
	fake.writeTagMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImagesReaderWriter) WriteTagCallCount() int {
	fake.writeTagMutex.RLock()
	defer fake.writeTagMutex.RUnlock()
	return len(fake.writeTagArgsForCall)
}

func (fake *FakeImagesReaderWriter) WriteTagCalls(stub func(name.Tag, remote.Taggable) error) {
	fake.writeTagMutex.Lock()
	defer fake.writeTagMutex.Unlock()
	fake.WriteTagStub = stub
}

func (fake *FakeImagesReaderWriter) WriteTagArgsForCall(i int) (name.Tag, remote.Taggable) {
	fake.writeTagMutex.RLock()
	defer fake.writeTagMutex.RUnlock()
	argsForCall := fake.writeTagArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImagesReaderWriter) WriteTagReturns(result1 error) {
	fake.writeTagMutex.Lock()
	defer fake.writeTagMutex.Unlock()
	fake.WriteTagStub = nil
	fake.writeTagReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) WriteTagReturnsOnCall(i int, result1 error) {
	fake.writeTagMutex.Lock()
	defer fake.writeTagMutex.Unlock()
	fake.WriteTagStub = nil
	if fake.writeTagReturnsOnCall == nil {
		fake.writeTagReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeTagReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImagesReaderWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.digestMutex.RLock()
	defer fake.digestMutex.RUnlock()
	fake.firstImageExistsMutex.RLock()
	defer fake.firstImageExistsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.multiWriteMutex.RLock()
	defer fake.multiWriteMutex.RUnlock()
	fake.writeImageMutex.RLock()
	defer fake.writeImageMutex.RUnlock()
	fake.writeIndexMutex.RLock()
	defer fake.writeIndexMutex.RUnlock()
	fake.writeTagMutex.RLock()
	defer fake.writeTagMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImagesReaderWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ imageset.ImagesReaderWriter = new(FakeImagesReaderWriter)
