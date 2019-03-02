// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/postal_address.proto

package postaladdress // import "github.com/Beeketing/go-genproto/googleapis/type/postaladdress"

import proto "github.com/Beeketing/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a postal address, e.g. for postal delivery or payments addresses.
// Given a postal address, a postal service can deliver items to a premise, P.O.
// Box or similar.
// It is not intended to model geographical locations (roads, towns,
// mountains).
//
// In typical usage an address would be created via user input or from importing
// existing data, depending on the type of process.
//
// Advice on address input / editing:
//  - Use an i18n-ready address widget such as
//    https://github.com/googlei18n/libaddressinput)
// - Users should not be presented with UI elements for input or editing of
//   fields outside countries where that field is used.
//
// For more guidance on how to use this schema, please see:
// https://support.google.com/business/answer/6397478
type PostalAddress struct {
	// The schema revision of the `PostalAddress`.
	// All new revisions **must** be backward compatible with old revisions.
	Revision int32 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. CLDR region code of the country/region of the address. This
	// is never inferred and it is up to the user to ensure the value is
	// correct. See http://cldr.unicode.org/ and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if
	// known). This is often the UI language of the input form or is expected
	// to match one of the languages used in the address' country/region, or their
	// transliterated equivalents.
	// This can affect formatting in certain countries, but is not critical
	// to the correctness of the data and will never affect any validation or
	// other non-formatting related operations.
	//
	// If this value is not known, it should be omitted (rather than specifying a
	// possibly incorrect default).
	//
	// Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require
	// postal codes to be present, but where they are used, they may trigger
	// additional validation with other parts of the address (e.g. state/zip
	// validation in the U.S.A.).
	PostalCode string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used
	// in most regions. Where it is used, the value is either a string like
	// "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	// alone, representing the "sector code" (Jamaica), "delivery area indicator"
	// (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode string `protobuf:"bytes,5,opt,name=sorting_code,json=sortingCode,proto3" json:"sorting_code,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal
	// addresses of a country or region.
	// For example, this can be a state, a province, an oblast, or a prefecture.
	// Specifically, for Spain this is the province and not the autonomous
	// community (e.g. "Barcelona" and not "Catalonia").
	// Many countries don't use an administrative area in postal addresses. E.g.
	// in Switzerland this should be left unpopulated.
	AdministrativeArea string `protobuf:"bytes,6,opt,name=administrative_area,json=administrativeArea,proto3" json:"administrative_area,omitempty"`
	// Optional. Generally refers to the city/town portion of the address.
	// Examples: US city, IT comune, UK post town.
	// In regions of the world where localities are not well defined or do not fit
	// into this structure well, leave locality empty and use address_lines.
	Locality string `protobuf:"bytes,7,opt,name=locality,proto3" json:"locality,omitempty"`
	// Optional. Sublocality of the address.
	// For example, this can be neighborhoods, boroughs, districts.
	Sublocality string `protobuf:"bytes,8,opt,name=sublocality,proto3" json:"sublocality,omitempty"`
	// Unstructured address lines describing the lower levels of an address.
	//
	// Because values in address_lines do not have type information and may
	// sometimes contain multiple values in a single field (e.g.
	// "Austin, TX"), it is important that the line order is clear. The order of
	// address lines should be "envelope order" for the country/region of the
	// address. In places where this can vary (e.g. Japan), address_language is
	// used to make it explicit (e.g. "ja" for large-to-small ordering and
	// "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	// an address can be selected based on the language.
	//
	// The minimum permitted structural representation of an address consists
	// of a region_code with all remaining information placed in the
	// address_lines. It would be possible to format such an address very
	// approximately without geocoding, but no semantic reasoning could be
	// made about any of the address components until it was at least
	// partially resolved.
	//
	// Creating an address only containing a region_code and address_lines, and
	// then geocoding is the recommended way to handle completely unstructured
	// addresses (as opposed to guessing which parts of the address should be
	// localities or administrative areas).
	AddressLines []string `protobuf:"bytes,9,rep,name=address_lines,json=addressLines,proto3" json:"address_lines,omitempty"`
	// Optional. The recipient at the address.
	// This field may, under certain circumstances, contain multiline information.
	// For example, it might contain "care of" information.
	Recipients []string `protobuf:"bytes,10,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Optional. The name of the organization at the address.
	Organization         string   `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostalAddress) Reset()         { *m = PostalAddress{} }
func (m *PostalAddress) String() string { return proto.CompactTextString(m) }
func (*PostalAddress) ProtoMessage()    {}
func (*PostalAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_postal_address_dc3389450c08c04a, []int{0}
}
func (m *PostalAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostalAddress.Unmarshal(m, b)
}
func (m *PostalAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostalAddress.Marshal(b, m, deterministic)
}
func (dst *PostalAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostalAddress.Merge(dst, src)
}
func (m *PostalAddress) XXX_Size() int {
	return xxx_messageInfo_PostalAddress.Size(m)
}
func (m *PostalAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PostalAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PostalAddress proto.InternalMessageInfo

func (m *PostalAddress) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *PostalAddress) GetRegionCode() string {
	if m != nil {
		return m.RegionCode
	}
	return ""
}

func (m *PostalAddress) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *PostalAddress) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *PostalAddress) GetSortingCode() string {
	if m != nil {
		return m.SortingCode
	}
	return ""
}

func (m *PostalAddress) GetAdministrativeArea() string {
	if m != nil {
		return m.AdministrativeArea
	}
	return ""
}

func (m *PostalAddress) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *PostalAddress) GetSublocality() string {
	if m != nil {
		return m.Sublocality
	}
	return ""
}

func (m *PostalAddress) GetAddressLines() []string {
	if m != nil {
		return m.AddressLines
	}
	return nil
}

func (m *PostalAddress) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *PostalAddress) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*PostalAddress)(nil), "google.type.PostalAddress")
}

func init() {
	proto.RegisterFile("google/type/postal_address.proto", fileDescriptor_postal_address_dc3389450c08c04a)
}

var fileDescriptor_postal_address_dc3389450c08c04a = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x31, 0x6f, 0xea, 0x30,
	0x10, 0xc7, 0x15, 0xf2, 0xe0, 0xc1, 0x05, 0xf4, 0x24, 0xbf, 0x25, 0xea, 0x50, 0x52, 0xba, 0x30,
	0x25, 0x43, 0xc7, 0x4e, 0x50, 0xa9, 0x5d, 0x3a, 0x44, 0xa8, 0x53, 0x97, 0xc8, 0x24, 0x27, 0xcb,
	0x52, 0xf0, 0x45, 0xb6, 0x41, 0xa2, 0xdf, 0xa1, 0x5f, 0xa2, 0x9f, 0xb4, 0xb2, 0x9d, 0xd2, 0x30,
	0xde, 0xef, 0x7e, 0x49, 0xee, 0xee, 0x1f, 0xc8, 0x04, 0x91, 0x68, 0xb1, 0xb0, 0xe7, 0x0e, 0x8b,
	0x8e, 0x8c, 0xe5, 0x6d, 0xc5, 0x9b, 0x46, 0xa3, 0x31, 0x79, 0xa7, 0xc9, 0x12, 0x4b, 0x82, 0x91,
	0x3b, 0x63, 0xf5, 0x19, 0xc3, 0xa2, 0xf4, 0xd6, 0x26, 0x48, 0xec, 0x06, 0xa6, 0x1a, 0x4f, 0xd2,
	0x48, 0x52, 0x69, 0x94, 0x45, 0xeb, 0xf1, 0xee, 0x52, 0xb3, 0x25, 0x24, 0x1a, 0x85, 0x24, 0x55,
	0xd5, 0xd4, 0x60, 0x3a, 0xca, 0xa2, 0xf5, 0x6c, 0x07, 0x01, 0x3d, 0x51, 0x83, 0xec, 0x1e, 0x16,
	0x2d, 0x57, 0xe2, 0xc8, 0x05, 0x06, 0x25, 0xf6, 0xca, 0xfc, 0x07, 0x7a, 0x69, 0x09, 0x49, 0x3f,
	0x98, 0x57, 0xfe, 0x84, 0xb7, 0x04, 0xe4, 0x85, 0x3b, 0x98, 0x1b, 0xd2, 0x56, 0x2a, 0x11, 0x8c,
	0xb1, 0x37, 0x92, 0x9e, 0x79, 0xa5, 0x80, 0xff, 0xbc, 0x39, 0x48, 0x25, 0x8d, 0xd5, 0xdc, 0xca,
	0x13, 0x56, 0x5c, 0x23, 0x4f, 0x27, 0xde, 0x64, 0xd7, 0xad, 0x8d, 0x46, 0xee, 0xd6, 0x6a, 0xa9,
	0xe6, 0xad, 0xb4, 0xe7, 0xf4, 0xaf, 0xb7, 0x2e, 0x35, 0xcb, 0x20, 0x31, 0xc7, 0xfd, 0xa5, 0x3d,
	0xed, 0x3f, 0xf7, 0x8b, 0xdc, 0x5e, 0xfd, 0x11, 0xab, 0x56, 0x2a, 0x34, 0xe9, 0x2c, 0x8b, 0xdd,
	0x5e, 0x3d, 0x7c, 0x75, 0x8c, 0xdd, 0x02, 0x68, 0xac, 0x65, 0x27, 0x51, 0x59, 0x93, 0x82, 0x37,
	0x06, 0x84, 0xad, 0x60, 0x4e, 0x5a, 0x70, 0x25, 0x3f, 0xb8, 0x75, 0xd7, 0x4d, 0xc2, 0x6d, 0x86,
	0x6c, 0x7b, 0x84, 0x7f, 0x35, 0x1d, 0xf2, 0x41, 0x44, 0x5b, 0x76, 0x95, 0x4f, 0xe9, 0x32, 0x2c,
	0xa3, 0xf7, 0xe7, 0x5e, 0x11, 0xe4, 0x6e, 0x9b, 0x93, 0x16, 0x85, 0x40, 0xe5, 0x13, 0x2e, 0x42,
	0x8b, 0x77, 0xd2, 0x0c, 0x7f, 0x83, 0x7e, 0xd6, 0xc7, 0xab, 0xea, 0x6b, 0x14, 0xbf, 0xbc, 0x95,
	0xfb, 0x89, 0x7f, 0xf0, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xda, 0x86, 0xd3, 0x22, 0x3e, 0x02,
	0x00, 0x00,
}
