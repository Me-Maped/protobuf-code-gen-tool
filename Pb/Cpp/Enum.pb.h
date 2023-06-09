// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Enum.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_Enum_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_Enum_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021012 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_Enum_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_Enum_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_Enum_2eproto;
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
namespace pb {

enum TestEnum1 : int {
  TEST_0 = 0,
  TEST_1 = 1,
  TestEnum1_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  TestEnum1_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool TestEnum1_IsValid(int value);
constexpr TestEnum1 TestEnum1_MIN = TEST_0;
constexpr TestEnum1 TestEnum1_MAX = TEST_1;
constexpr int TestEnum1_ARRAYSIZE = TestEnum1_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* TestEnum1_descriptor();
template<typename T>
inline const std::string& TestEnum1_Name(T enum_t_value) {
  static_assert(::std::is_same<T, TestEnum1>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function TestEnum1_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    TestEnum1_descriptor(), enum_t_value);
}
inline bool TestEnum1_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, TestEnum1* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<TestEnum1>(
    TestEnum1_descriptor(), name, value);
}
enum TestEnum2 : int {
  TEST_3 = 0,
  TEST_4 = 4,
  TestEnum2_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  TestEnum2_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool TestEnum2_IsValid(int value);
constexpr TestEnum2 TestEnum2_MIN = TEST_3;
constexpr TestEnum2 TestEnum2_MAX = TEST_4;
constexpr int TestEnum2_ARRAYSIZE = TestEnum2_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* TestEnum2_descriptor();
template<typename T>
inline const std::string& TestEnum2_Name(T enum_t_value) {
  static_assert(::std::is_same<T, TestEnum2>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function TestEnum2_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    TestEnum2_descriptor(), enum_t_value);
}
inline bool TestEnum2_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, TestEnum2* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<TestEnum2>(
    TestEnum2_descriptor(), name, value);
}
// ===================================================================


// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace pb

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::pb::TestEnum1> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::pb::TestEnum1>() {
  return ::pb::TestEnum1_descriptor();
}
template <> struct is_proto_enum< ::pb::TestEnum2> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::pb::TestEnum2>() {
  return ::pb::TestEnum2_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_Enum_2eproto
