// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.rpc/v1/federated_gloo_resources.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/mitchellh/hashstructure"
	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *FederatedUpstream) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.FederatedUpstream")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedUpstreamGroup) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.FederatedUpstreamGroup")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *FederatedSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.FederatedSettings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSpec()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetSpec(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedUpstreamsRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedUpstreamsRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedUpstreamsResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedUpstreamsResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedUpstreams() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedUpstreamYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedUpstreamYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetFederatedUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedUpstreamYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedUpstreamYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedUpstreamGroupsRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedUpstreamGroupsRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedUpstreamGroupsResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedUpstreamGroupsResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedUpstreamGroups() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedUpstreamGroupYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedUpstreamGroupYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedUpstreamGroupRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetFederatedUpstreamGroupRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedUpstreamGroupYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedUpstreamGroupYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedSettingsRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedSettingsRequest")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ListFederatedSettingsResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.ListFederatedSettingsResponse")); err != nil {
		return 0, err
	}

	for _, v := range m.GetFederatedSettings() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if val, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedSettingsYamlRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedSettingsYamlRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetFederatedSettingsRef()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetFederatedSettingsRef(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *GetFederatedSettingsYamlResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("fed.rpc.solo.io.github.com/solo-io/solo-apis/pkg/api/fed.rpc/v1.GetFederatedSettingsYamlResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetYamlData()).(safe_hasher.SafeHasher); ok {
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if val, err := hashstructure.Hash(m.GetYamlData(), nil); err != nil {
			return 0, err
		} else {
			if err := binary.Write(hasher, binary.LittleEndian, val); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
