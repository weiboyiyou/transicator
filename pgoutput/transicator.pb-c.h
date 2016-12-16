/*
Copyright 2016 The Transicator Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: transicator.proto */

#ifndef PROTOBUF_C_transicator_2eproto__INCLUDED
#define PROTOBUF_C_transicator_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1002001 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _Common__ValuePb Common__ValuePb;
typedef struct _Common__ColumnPb Common__ColumnPb;
typedef struct _Common__ChangePb Common__ChangePb;
typedef struct _Common__ChangeListPb Common__ChangeListPb;
typedef struct _Common__SnapshotHeaderPb Common__SnapshotHeaderPb;
typedef struct _Common__TableHeaderPb Common__TableHeaderPb;
typedef struct _Common__RowPb Common__RowPb;
typedef struct _Common__StreamMessagePb Common__StreamMessagePb;


/* --- enums --- */


/* --- messages --- */

typedef enum {
  COMMON__VALUE_PB__VALUE__NOT_SET = 0,
  COMMON__VALUE_PB__VALUE_STRING = 1,
  COMMON__VALUE_PB__VALUE_INT = 2,
  COMMON__VALUE_PB__VALUE_UINT = 3,
  COMMON__VALUE_PB__VALUE_DOUBLE = 4,
  COMMON__VALUE_PB__VALUE_BYTES = 5,
  COMMON__VALUE_PB__VALUE_BOOL = 6,
  COMMON__VALUE_PB__VALUE_TIMESTAMP = 7,
} Common__ValuePb__ValueCase;

struct  _Common__ValuePb
{
  ProtobufCMessage base;
  Common__ValuePb__ValueCase value_case;
  union {
    char *string;
    int64_t int_;
    uint64_t uint;
    double double_;
    ProtobufCBinaryData bytes;
    protobuf_c_boolean bool_;
    int64_t timestamp;
  };
};
#define COMMON__VALUE_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__value_pb__descriptor) \
    , COMMON__VALUE_PB__VALUE__NOT_SET, {0} }


struct  _Common__ColumnPb
{
  ProtobufCMessage base;
  char *name;
  Common__ValuePb *value;
  protobuf_c_boolean has_type;
  int32_t type;
};
#define COMMON__COLUMN_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__column_pb__descriptor) \
    , NULL, NULL, 0,0 }


struct  _Common__ChangePb
{
  ProtobufCMessage base;
  /*
   * operation -- May be 1, 2, 3 for insert, update, or delete
   */
  int32_t operation;
  /*
   * Table name, in "schema.table" format
   */
  char *table;
  /*
   * Sequence number is generated by the change server by combining
   * the commitSequence and commitIndex
   */
  char *sequence;
  /*
   * LSN when the transaction was committed. Primary sort order.
   */
  protobuf_c_boolean has_commitsequence;
  uint64_t commitsequence;
  /*
   * LSN when the change was made. Mostly of academic interest.
   */
  protobuf_c_boolean has_changesequence;
  uint64_t changesequence;
  /*
   * Order in which the change was made inside the transaction.
   * Secondary sort order.
   */
  protobuf_c_boolean has_commitindex;
  uint32_t commitindex;
  /*
   * 32-bit Postgres transaction ID. Not used by newer code.
   * Rolls over.
   */
  protobuf_c_boolean has_transactionid;
  uint32_t transactionid;
  /*
   * Row added to an insert, or new row for an update.
   */
  size_t n_newcolumns;
  Common__ColumnPb **newcolumns;
  /*
   * Row removed on a delete, or old row for an update.
   */
  size_t n_oldcolumns;
  Common__ColumnPb **oldcolumns;
  /*
   * Time when the record was inserted, in seconds since the
   * Unix epoch.
   */
  protobuf_c_boolean has_timestamp;
  int64_t timestamp;
  /*
   * 64-bit Postgres transaction id, including the epoch. Will not
   * roll over.
   */
  protobuf_c_boolean has_transactionidepoch;
  uint64_t transactionidepoch;
};
#define COMMON__CHANGE_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__change_pb__descriptor) \
    , 0, NULL, NULL, 0,0, 0,0, 0,0, 0,0, 0,NULL, 0,NULL, 0,0, 0,0 }


struct  _Common__ChangeListPb
{
  ProtobufCMessage base;
  char *lastsequence;
  char *firstsequence;
  size_t n_changes;
  Common__ChangePb **changes;
};
#define COMMON__CHANGE_LIST_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__change_list_pb__descriptor) \
    , NULL, NULL, 0,NULL }


struct  _Common__SnapshotHeaderPb
{
  ProtobufCMessage base;
  char *timestamp;
  char *snapshot;
};
#define COMMON__SNAPSHOT_HEADER_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__snapshot_header_pb__descriptor) \
    , NULL, NULL }


struct  _Common__TableHeaderPb
{
  ProtobufCMessage base;
  char *name;
  size_t n_columns;
  Common__ColumnPb **columns;
};
#define COMMON__TABLE_HEADER_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__table_header_pb__descriptor) \
    , NULL, 0,NULL }


struct  _Common__RowPb
{
  ProtobufCMessage base;
  size_t n_values;
  Common__ValuePb **values;
};
#define COMMON__ROW_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__row_pb__descriptor) \
    , 0,NULL }


typedef enum {
  COMMON__STREAM_MESSAGE_PB__MESSAGE__NOT_SET = 0,
  COMMON__STREAM_MESSAGE_PB__MESSAGE_TABLE = 1,
  COMMON__STREAM_MESSAGE_PB__MESSAGE_ROW = 2,
} Common__StreamMessagePb__MessageCase;

struct  _Common__StreamMessagePb
{
  ProtobufCMessage base;
  Common__StreamMessagePb__MessageCase message_case;
  union {
    Common__TableHeaderPb *table;
    Common__RowPb *row;
  };
};
#define COMMON__STREAM_MESSAGE_PB__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&common__stream_message_pb__descriptor) \
    , COMMON__STREAM_MESSAGE_PB__MESSAGE__NOT_SET, {0} }


/* Common__ValuePb methods */
void   common__value_pb__init
                     (Common__ValuePb         *message);
size_t common__value_pb__get_packed_size
                     (const Common__ValuePb   *message);
size_t common__value_pb__pack
                     (const Common__ValuePb   *message,
                      uint8_t             *out);
size_t common__value_pb__pack_to_buffer
                     (const Common__ValuePb   *message,
                      ProtobufCBuffer     *buffer);
Common__ValuePb *
       common__value_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__value_pb__free_unpacked
                     (Common__ValuePb *message,
                      ProtobufCAllocator *allocator);
/* Common__ColumnPb methods */
void   common__column_pb__init
                     (Common__ColumnPb         *message);
size_t common__column_pb__get_packed_size
                     (const Common__ColumnPb   *message);
size_t common__column_pb__pack
                     (const Common__ColumnPb   *message,
                      uint8_t             *out);
size_t common__column_pb__pack_to_buffer
                     (const Common__ColumnPb   *message,
                      ProtobufCBuffer     *buffer);
Common__ColumnPb *
       common__column_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__column_pb__free_unpacked
                     (Common__ColumnPb *message,
                      ProtobufCAllocator *allocator);
/* Common__ChangePb methods */
void   common__change_pb__init
                     (Common__ChangePb         *message);
size_t common__change_pb__get_packed_size
                     (const Common__ChangePb   *message);
size_t common__change_pb__pack
                     (const Common__ChangePb   *message,
                      uint8_t             *out);
size_t common__change_pb__pack_to_buffer
                     (const Common__ChangePb   *message,
                      ProtobufCBuffer     *buffer);
Common__ChangePb *
       common__change_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__change_pb__free_unpacked
                     (Common__ChangePb *message,
                      ProtobufCAllocator *allocator);
/* Common__ChangeListPb methods */
void   common__change_list_pb__init
                     (Common__ChangeListPb         *message);
size_t common__change_list_pb__get_packed_size
                     (const Common__ChangeListPb   *message);
size_t common__change_list_pb__pack
                     (const Common__ChangeListPb   *message,
                      uint8_t             *out);
size_t common__change_list_pb__pack_to_buffer
                     (const Common__ChangeListPb   *message,
                      ProtobufCBuffer     *buffer);
Common__ChangeListPb *
       common__change_list_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__change_list_pb__free_unpacked
                     (Common__ChangeListPb *message,
                      ProtobufCAllocator *allocator);
/* Common__SnapshotHeaderPb methods */
void   common__snapshot_header_pb__init
                     (Common__SnapshotHeaderPb         *message);
size_t common__snapshot_header_pb__get_packed_size
                     (const Common__SnapshotHeaderPb   *message);
size_t common__snapshot_header_pb__pack
                     (const Common__SnapshotHeaderPb   *message,
                      uint8_t             *out);
size_t common__snapshot_header_pb__pack_to_buffer
                     (const Common__SnapshotHeaderPb   *message,
                      ProtobufCBuffer     *buffer);
Common__SnapshotHeaderPb *
       common__snapshot_header_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__snapshot_header_pb__free_unpacked
                     (Common__SnapshotHeaderPb *message,
                      ProtobufCAllocator *allocator);
/* Common__TableHeaderPb methods */
void   common__table_header_pb__init
                     (Common__TableHeaderPb         *message);
size_t common__table_header_pb__get_packed_size
                     (const Common__TableHeaderPb   *message);
size_t common__table_header_pb__pack
                     (const Common__TableHeaderPb   *message,
                      uint8_t             *out);
size_t common__table_header_pb__pack_to_buffer
                     (const Common__TableHeaderPb   *message,
                      ProtobufCBuffer     *buffer);
Common__TableHeaderPb *
       common__table_header_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__table_header_pb__free_unpacked
                     (Common__TableHeaderPb *message,
                      ProtobufCAllocator *allocator);
/* Common__RowPb methods */
void   common__row_pb__init
                     (Common__RowPb         *message);
size_t common__row_pb__get_packed_size
                     (const Common__RowPb   *message);
size_t common__row_pb__pack
                     (const Common__RowPb   *message,
                      uint8_t             *out);
size_t common__row_pb__pack_to_buffer
                     (const Common__RowPb   *message,
                      ProtobufCBuffer     *buffer);
Common__RowPb *
       common__row_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__row_pb__free_unpacked
                     (Common__RowPb *message,
                      ProtobufCAllocator *allocator);
/* Common__StreamMessagePb methods */
void   common__stream_message_pb__init
                     (Common__StreamMessagePb         *message);
size_t common__stream_message_pb__get_packed_size
                     (const Common__StreamMessagePb   *message);
size_t common__stream_message_pb__pack
                     (const Common__StreamMessagePb   *message,
                      uint8_t             *out);
size_t common__stream_message_pb__pack_to_buffer
                     (const Common__StreamMessagePb   *message,
                      ProtobufCBuffer     *buffer);
Common__StreamMessagePb *
       common__stream_message_pb__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   common__stream_message_pb__free_unpacked
                     (Common__StreamMessagePb *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*Common__ValuePb_Closure)
                 (const Common__ValuePb *message,
                  void *closure_data);
typedef void (*Common__ColumnPb_Closure)
                 (const Common__ColumnPb *message,
                  void *closure_data);
typedef void (*Common__ChangePb_Closure)
                 (const Common__ChangePb *message,
                  void *closure_data);
typedef void (*Common__ChangeListPb_Closure)
                 (const Common__ChangeListPb *message,
                  void *closure_data);
typedef void (*Common__SnapshotHeaderPb_Closure)
                 (const Common__SnapshotHeaderPb *message,
                  void *closure_data);
typedef void (*Common__TableHeaderPb_Closure)
                 (const Common__TableHeaderPb *message,
                  void *closure_data);
typedef void (*Common__RowPb_Closure)
                 (const Common__RowPb *message,
                  void *closure_data);
typedef void (*Common__StreamMessagePb_Closure)
                 (const Common__StreamMessagePb *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor common__value_pb__descriptor;
extern const ProtobufCMessageDescriptor common__column_pb__descriptor;
extern const ProtobufCMessageDescriptor common__change_pb__descriptor;
extern const ProtobufCMessageDescriptor common__change_list_pb__descriptor;
extern const ProtobufCMessageDescriptor common__snapshot_header_pb__descriptor;
extern const ProtobufCMessageDescriptor common__table_header_pb__descriptor;
extern const ProtobufCMessageDescriptor common__row_pb__descriptor;
extern const ProtobufCMessageDescriptor common__stream_message_pb__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_transicator_2eproto__INCLUDED */
