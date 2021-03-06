// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *GetStoredApplicationUpRequest) SetFields(src *GetStoredApplicationUpRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIDs == nil) && dst.ApplicationIDs == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIDs
				}
				if dst.ApplicationIDs != nil {
					newDst = dst.ApplicationIDs
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIDs = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIDs = src.ApplicationIDs
				} else {
					dst.ApplicationIDs = nil
				}
			}
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIDs == nil) && dst.EndDeviceIDs == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIDs
				}
				if dst.EndDeviceIDs != nil {
					newDst = dst.EndDeviceIDs
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIDs = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIDs = src.EndDeviceIDs
				} else {
					dst.EndDeviceIDs = nil
				}
			}
		case "type":
			if len(subs) > 0 {
				return fmt.Errorf("'type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Type = src.Type
			} else {
				var zero string
				dst.Type = zero
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				dst.Limit = nil
			}
		case "after":
			if len(subs) > 0 {
				return fmt.Errorf("'after' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.After = src.After
			} else {
				dst.After = nil
			}
		case "before":
			if len(subs) > 0 {
				return fmt.Errorf("'before' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Before = src.Before
			} else {
				dst.Before = nil
			}
		case "f_port":
			if len(subs) > 0 {
				return fmt.Errorf("'f_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FPort = src.FPort
			} else {
				dst.FPort = nil
			}
		case "order":
			if len(subs) > 0 {
				return fmt.Errorf("'order' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Order = src.Order
			} else {
				var zero string
				dst.Order = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
