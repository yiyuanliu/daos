/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: DunsAttribute.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "DunsAttribute.pb-c.h"
void   uns__daos_ace__init
                     (Uns__DaosAce         *message)
{
  static const Uns__DaosAce init_value = UNS__DAOS_ACE__INIT;
  *message = init_value;
}
size_t uns__daos_ace__get_packed_size
                     (const Uns__DaosAce *message)
{
  assert(message->base.descriptor == &uns__daos_ace__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t uns__daos_ace__pack
                     (const Uns__DaosAce *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &uns__daos_ace__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t uns__daos_ace__pack_to_buffer
                     (const Uns__DaosAce *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &uns__daos_ace__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Uns__DaosAce *
       uns__daos_ace__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Uns__DaosAce *)
     protobuf_c_message_unpack (&uns__daos_ace__descriptor,
                                allocator, len, data);
}
void   uns__daos_ace__free_unpacked
                     (Uns__DaosAce *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &uns__daos_ace__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   uns__daos_acl__init
                     (Uns__DaosAcl         *message)
{
  static const Uns__DaosAcl init_value = UNS__DAOS_ACL__INIT;
  *message = init_value;
}
size_t uns__daos_acl__get_packed_size
                     (const Uns__DaosAcl *message)
{
  assert(message->base.descriptor == &uns__daos_acl__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t uns__daos_acl__pack
                     (const Uns__DaosAcl *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &uns__daos_acl__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t uns__daos_acl__pack_to_buffer
                     (const Uns__DaosAcl *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &uns__daos_acl__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Uns__DaosAcl *
       uns__daos_acl__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Uns__DaosAcl *)
     protobuf_c_message_unpack (&uns__daos_acl__descriptor,
                                allocator, len, data);
}
void   uns__daos_acl__free_unpacked
                     (Uns__DaosAcl *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &uns__daos_acl__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   uns__entry__init
                     (Uns__Entry         *message)
{
  static const Uns__Entry init_value = UNS__ENTRY__INIT;
  *message = init_value;
}
size_t uns__entry__get_packed_size
                     (const Uns__Entry *message)
{
  assert(message->base.descriptor == &uns__entry__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t uns__entry__pack
                     (const Uns__Entry *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &uns__entry__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t uns__entry__pack_to_buffer
                     (const Uns__Entry *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &uns__entry__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Uns__Entry *
       uns__entry__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Uns__Entry *)
     protobuf_c_message_unpack (&uns__entry__descriptor,
                                allocator, len, data);
}
void   uns__entry__free_unpacked
                     (Uns__Entry *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &uns__entry__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   uns__properties__init
                     (Uns__Properties         *message)
{
  static const Uns__Properties init_value = UNS__PROPERTIES__INIT;
  *message = init_value;
}
size_t uns__properties__get_packed_size
                     (const Uns__Properties *message)
{
  assert(message->base.descriptor == &uns__properties__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t uns__properties__pack
                     (const Uns__Properties *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &uns__properties__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t uns__properties__pack_to_buffer
                     (const Uns__Properties *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &uns__properties__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Uns__Properties *
       uns__properties__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Uns__Properties *)
     protobuf_c_message_unpack (&uns__properties__descriptor,
                                allocator, len, data);
}
void   uns__properties__free_unpacked
                     (Uns__Properties *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &uns__properties__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   uns__duns_attribute__init
                     (Uns__DunsAttribute         *message)
{
  static const Uns__DunsAttribute init_value = UNS__DUNS_ATTRIBUTE__INIT;
  *message = init_value;
}
size_t uns__duns_attribute__get_packed_size
                     (const Uns__DunsAttribute *message)
{
  assert(message->base.descriptor == &uns__duns_attribute__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t uns__duns_attribute__pack
                     (const Uns__DunsAttribute *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &uns__duns_attribute__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t uns__duns_attribute__pack_to_buffer
                     (const Uns__DunsAttribute *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &uns__duns_attribute__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Uns__DunsAttribute *
       uns__duns_attribute__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Uns__DunsAttribute *)
     protobuf_c_message_unpack (&uns__duns_attribute__descriptor,
                                allocator, len, data);
}
void   uns__duns_attribute__free_unpacked
                     (Uns__DunsAttribute *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &uns__duns_attribute__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor uns__daos_ace__field_descriptors[9] =
{
  {
    "access_types",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, access_types),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "principal_type",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, principal_type),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "principal_len",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, principal_len),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "access_flags",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, access_flags),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "reserved",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, reserved),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "allow_perms",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, allow_perms),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "audit_perms",
    7,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, audit_perms),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "alarm_perms",
    8,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, alarm_perms),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "principal",
    9,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAce, principal),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned uns__daos_ace__field_indices_by_name[] = {
  3,   /* field[3] = access_flags */
  0,   /* field[0] = access_types */
  7,   /* field[7] = alarm_perms */
  5,   /* field[5] = allow_perms */
  6,   /* field[6] = audit_perms */
  8,   /* field[8] = principal */
  2,   /* field[2] = principal_len */
  1,   /* field[1] = principal_type */
  4,   /* field[4] = reserved */
};
static const ProtobufCIntRange uns__daos_ace__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 9 }
};
const ProtobufCMessageDescriptor uns__daos_ace__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "uns.DaosAce",
  "DaosAce",
  "Uns__DaosAce",
  "uns",
  sizeof(Uns__DaosAce),
  9,
  uns__daos_ace__field_descriptors,
  uns__daos_ace__field_indices_by_name,
  1,  uns__daos_ace__number_ranges,
  (ProtobufCMessageInit) uns__daos_ace__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor uns__daos_acl__field_descriptors[3] =
{
  {
    "ver",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAcl, ver),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "reserv",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__DaosAcl, reserv),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "aces",
    4,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Uns__DaosAcl, n_aces),
    offsetof(Uns__DaosAcl, aces),
    &uns__daos_ace__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned uns__daos_acl__field_indices_by_name[] = {
  2,   /* field[2] = aces */
  1,   /* field[1] = reserv */
  0,   /* field[0] = ver */
};
static const ProtobufCIntRange uns__daos_acl__number_ranges[2 + 1] =
{
  { 1, 0 },
  { 4, 2 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor uns__daos_acl__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "uns.DaosAcl",
  "DaosAcl",
  "Uns__DaosAcl",
  "uns",
  sizeof(Uns__DaosAcl),
  3,
  uns__daos_acl__field_descriptors,
  uns__daos_acl__field_indices_by_name,
  2,  uns__daos_acl__number_ranges,
  (ProtobufCMessageInit) uns__daos_acl__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor uns__entry__field_descriptors[5] =
{
  {
    "type",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Uns__Entry, type),
    &uns__prop_type__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "reserved",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__Entry, reserved),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "val",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    offsetof(Uns__Entry, value_case),
    offsetof(Uns__Entry, val),
    NULL,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "str",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    offsetof(Uns__Entry, value_case),
    offsetof(Uns__Entry, str),
    NULL,
    &protobuf_c_empty_string,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "pval",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Uns__Entry, value_case),
    offsetof(Uns__Entry, pval),
    &uns__daos_acl__descriptor,
    NULL,
    0 | PROTOBUF_C_FIELD_FLAG_ONEOF,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned uns__entry__field_indices_by_name[] = {
  4,   /* field[4] = pval */
  1,   /* field[1] = reserved */
  3,   /* field[3] = str */
  0,   /* field[0] = type */
  2,   /* field[2] = val */
};
static const ProtobufCIntRange uns__entry__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 5 }
};
const ProtobufCMessageDescriptor uns__entry__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "uns.Entry",
  "Entry",
  "Uns__Entry",
  "uns",
  sizeof(Uns__Entry),
  5,
  uns__entry__field_descriptors,
  uns__entry__field_indices_by_name,
  1,  uns__entry__number_ranges,
  (ProtobufCMessageInit) uns__entry__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor uns__properties__field_descriptors[2] =
{
  {
    "reserved",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Uns__Properties, reserved),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "entries",
    2,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_MESSAGE,
    offsetof(Uns__Properties, n_entries),
    offsetof(Uns__Properties, entries),
    &uns__entry__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned uns__properties__field_indices_by_name[] = {
  1,   /* field[1] = entries */
  0,   /* field[0] = reserved */
};
static const ProtobufCIntRange uns__properties__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor uns__properties__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "uns.Properties",
  "Properties",
  "Uns__Properties",
  "uns",
  sizeof(Uns__Properties),
  2,
  uns__properties__field_descriptors,
  uns__properties__field_indices_by_name,
  1,  uns__properties__number_ranges,
  (ProtobufCMessageInit) uns__properties__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor uns__duns_attribute__field_descriptors[9] =
{
  {
    "puuid",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, puuid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "cuuid",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, cuuid),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "layout_type",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, layout_type),
    &uns__layout__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "object_type",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, object_type),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "chunk_size",
    5,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, chunk_size),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "rel_path",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, rel_path),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "on_lustre",
    7,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_BOOL,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, on_lustre),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "properties",
    8,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, properties),
    &uns__properties__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "no_prefix",
    9,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_BOOL,
    0,   /* quantifier_offset */
    offsetof(Uns__DunsAttribute, no_prefix),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned uns__duns_attribute__field_indices_by_name[] = {
  4,   /* field[4] = chunk_size */
  1,   /* field[1] = cuuid */
  2,   /* field[2] = layout_type */
  8,   /* field[8] = no_prefix */
  3,   /* field[3] = object_type */
  6,   /* field[6] = on_lustre */
  7,   /* field[7] = properties */
  0,   /* field[0] = puuid */
  5,   /* field[5] = rel_path */
};
static const ProtobufCIntRange uns__duns_attribute__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 9 }
};
const ProtobufCMessageDescriptor uns__duns_attribute__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "uns.DunsAttribute",
  "DunsAttribute",
  "Uns__DunsAttribute",
  "uns",
  sizeof(Uns__DunsAttribute),
  9,
  uns__duns_attribute__field_descriptors,
  uns__duns_attribute__field_indices_by_name,
  1,  uns__duns_attribute__number_ranges,
  (ProtobufCMessageInit) uns__duns_attribute__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue uns__prop_type__enum_values_by_number[26] =
{
  { "DAOS_PROP_PO_MIN", "UNS__PROP_TYPE__DAOS_PROP_PO_MIN", 0 },
  { "DAOS_PROP_PO_LABEL", "UNS__PROP_TYPE__DAOS_PROP_PO_LABEL", 1 },
  { "DAOS_PROP_PO_ACL", "UNS__PROP_TYPE__DAOS_PROP_PO_ACL", 2 },
  { "DAOS_PROP_PO_SPACE_RB", "UNS__PROP_TYPE__DAOS_PROP_PO_SPACE_RB", 3 },
  { "DAOS_PROP_PO_SELF_HEAL", "UNS__PROP_TYPE__DAOS_PROP_PO_SELF_HEAL", 4 },
  { "DAOS_PROP_PO_RECLAIM", "UNS__PROP_TYPE__DAOS_PROP_PO_RECLAIM", 5 },
  { "DAOS_PROP_PO_OWNER", "UNS__PROP_TYPE__DAOS_PROP_PO_OWNER", 6 },
  { "DAOS_PROP_PO_OWNER_GROUP", "UNS__PROP_TYPE__DAOS_PROP_PO_OWNER_GROUP", 7 },
  { "DAOS_PROP_PO_SVC_LIST", "UNS__PROP_TYPE__DAOS_PROP_PO_SVC_LIST", 8 },
  { "DAOS_PROP_PO_MAX", "UNS__PROP_TYPE__DAOS_PROP_PO_MAX", 9 },
  { "DAOS_PROP_CO_MIN", "UNS__PROP_TYPE__DAOS_PROP_CO_MIN", 4096 },
  { "DAOS_PROP_CO_LABEL", "UNS__PROP_TYPE__DAOS_PROP_CO_LABEL", 4097 },
  { "DAOS_PROP_CO_LAYOUT_TYPE", "UNS__PROP_TYPE__DAOS_PROP_CO_LAYOUT_TYPE", 4098 },
  { "DAOS_PROP_CO_LAYOUT_VER", "UNS__PROP_TYPE__DAOS_PROP_CO_LAYOUT_VER", 4099 },
  { "DAOS_PROP_CO_CSUM", "UNS__PROP_TYPE__DAOS_PROP_CO_CSUM", 4100 },
  { "DAOS_PROP_CO_CSUM_CHUNK_SIZE", "UNS__PROP_TYPE__DAOS_PROP_CO_CSUM_CHUNK_SIZE", 4101 },
  { "DAOS_PROP_CO_CSUM_SERVER_VERIFY", "UNS__PROP_TYPE__DAOS_PROP_CO_CSUM_SERVER_VERIFY", 4102 },
  { "DAOS_PROP_CO_REDUN_FAC", "UNS__PROP_TYPE__DAOS_PROP_CO_REDUN_FAC", 4103 },
  { "DAOS_PROP_CO_REDUN_LVL", "UNS__PROP_TYPE__DAOS_PROP_CO_REDUN_LVL", 4104 },
  { "DAOS_PROP_CO_SNAPSHOT_MAX", "UNS__PROP_TYPE__DAOS_PROP_CO_SNAPSHOT_MAX", 4105 },
  { "DAOS_PROP_CO_ACL", "UNS__PROP_TYPE__DAOS_PROP_CO_ACL", 4106 },
  { "DAOS_PROP_CO_COMPRESS", "UNS__PROP_TYPE__DAOS_PROP_CO_COMPRESS", 4107 },
  { "DAOS_PROP_CO_ENCRYPT", "UNS__PROP_TYPE__DAOS_PROP_CO_ENCRYPT", 4108 },
  { "DAOS_PROP_CO_OWNER", "UNS__PROP_TYPE__DAOS_PROP_CO_OWNER", 4109 },
  { "DAOS_PROP_CO_OWNER_GROUP", "UNS__PROP_TYPE__DAOS_PROP_CO_OWNER_GROUP", 4110 },
  { "DAOS_PROP_CO_MAX", "UNS__PROP_TYPE__DAOS_PROP_CO_MAX", 4111 },
};
static const ProtobufCIntRange uns__prop_type__value_ranges[] = {
{0, 0},{4096, 10},{0, 26}
};
static const ProtobufCEnumValueIndex uns__prop_type__enum_values_by_name[26] =
{
  { "DAOS_PROP_CO_ACL", 20 },
  { "DAOS_PROP_CO_COMPRESS", 21 },
  { "DAOS_PROP_CO_CSUM", 14 },
  { "DAOS_PROP_CO_CSUM_CHUNK_SIZE", 15 },
  { "DAOS_PROP_CO_CSUM_SERVER_VERIFY", 16 },
  { "DAOS_PROP_CO_ENCRYPT", 22 },
  { "DAOS_PROP_CO_LABEL", 11 },
  { "DAOS_PROP_CO_LAYOUT_TYPE", 12 },
  { "DAOS_PROP_CO_LAYOUT_VER", 13 },
  { "DAOS_PROP_CO_MAX", 25 },
  { "DAOS_PROP_CO_MIN", 10 },
  { "DAOS_PROP_CO_OWNER", 23 },
  { "DAOS_PROP_CO_OWNER_GROUP", 24 },
  { "DAOS_PROP_CO_REDUN_FAC", 17 },
  { "DAOS_PROP_CO_REDUN_LVL", 18 },
  { "DAOS_PROP_CO_SNAPSHOT_MAX", 19 },
  { "DAOS_PROP_PO_ACL", 2 },
  { "DAOS_PROP_PO_LABEL", 1 },
  { "DAOS_PROP_PO_MAX", 9 },
  { "DAOS_PROP_PO_MIN", 0 },
  { "DAOS_PROP_PO_OWNER", 6 },
  { "DAOS_PROP_PO_OWNER_GROUP", 7 },
  { "DAOS_PROP_PO_RECLAIM", 5 },
  { "DAOS_PROP_PO_SELF_HEAL", 4 },
  { "DAOS_PROP_PO_SPACE_RB", 3 },
  { "DAOS_PROP_PO_SVC_LIST", 8 },
};
const ProtobufCEnumDescriptor uns__prop_type__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "uns.PropType",
  "PropType",
  "Uns__PropType",
  "uns",
  26,
  uns__prop_type__enum_values_by_number,
  26,
  uns__prop_type__enum_values_by_name,
  2,
  uns__prop_type__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
static const ProtobufCEnumValue uns__layout__enum_values_by_number[3] =
{
  { "UNKNOWN", "UNS__LAYOUT__UNKNOWN", 0 },
  { "POSIX", "UNS__LAYOUT__POSIX", 1 },
  { "HDF5", "UNS__LAYOUT__HDF5", 2 },
};
static const ProtobufCIntRange uns__layout__value_ranges[] = {
{0, 0},{0, 3}
};
static const ProtobufCEnumValueIndex uns__layout__enum_values_by_name[3] =
{
  { "HDF5", 2 },
  { "POSIX", 1 },
  { "UNKNOWN", 0 },
};
const ProtobufCEnumDescriptor uns__layout__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "uns.Layout",
  "Layout",
  "Uns__Layout",
  "uns",
  3,
  uns__layout__enum_values_by_number,
  3,
  uns__layout__enum_values_by_name,
  1,
  uns__layout__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};