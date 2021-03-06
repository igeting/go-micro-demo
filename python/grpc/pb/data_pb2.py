# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: data.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='data.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\ndata.proto\x12\x02pb\"\xa3\x01\n\ractionRequest\x12\x0c\n\x04text\x18\x01 \x01(\t\x12(\n\x06\x63orpus\x18\x04 \x01(\x0e\x32\x18.pb.actionRequest.Corpus\"Z\n\x06\x43orpus\x12\r\n\tUNIVERSAL\x10\x00\x12\x07\n\x03WEB\x10\x01\x12\n\n\x06IMAGES\x10\x02\x12\t\n\x05LOCAL\x10\x03\x12\x08\n\x04NEWS\x10\x04\x12\x0c\n\x08PRODUCTS\x10\x05\x12\t\n\x05VIDEO\x10\x06\"\x8e\x01\n\x0e\x61\x63tionResponse\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x0b\n\x03\x61ge\x18\x02 \x01(\x05\x12)\n\x06result\x18\x03 \x03(\x0b\x32\x19.pb.actionResponse.Result\x1a\x36\n\x06Result\x12\x0b\n\x03url\x18\x01 \x01(\t\x12\r\n\x05title\x18\x02 \x01(\t\x12\x10\n\x08snippets\x18\x03 \x03(\t2A\n\nFormatData\x12\x33\n\x08\x44oFormat\x12\x11.pb.actionRequest\x1a\x12.pb.actionResponse\"\x00\x62\x06proto3'
)



_ACTIONREQUEST_CORPUS = _descriptor.EnumDescriptor(
  name='Corpus',
  full_name='pb.actionRequest.Corpus',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNIVERSAL', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='WEB', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='IMAGES', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='LOCAL', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NEWS', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PRODUCTS', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='VIDEO', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=92,
  serialized_end=182,
)
_sym_db.RegisterEnumDescriptor(_ACTIONREQUEST_CORPUS)


_ACTIONREQUEST = _descriptor.Descriptor(
  name='actionRequest',
  full_name='pb.actionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='pb.actionRequest.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='corpus', full_name='pb.actionRequest.corpus', index=1,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _ACTIONREQUEST_CORPUS,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=19,
  serialized_end=182,
)


_ACTIONRESPONSE_RESULT = _descriptor.Descriptor(
  name='Result',
  full_name='pb.actionResponse.Result',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='url', full_name='pb.actionResponse.Result.url', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='title', full_name='pb.actionResponse.Result.title', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='snippets', full_name='pb.actionResponse.Result.snippets', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=273,
  serialized_end=327,
)

_ACTIONRESPONSE = _descriptor.Descriptor(
  name='actionResponse',
  full_name='pb.actionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='text', full_name='pb.actionResponse.text', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='age', full_name='pb.actionResponse.age', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result', full_name='pb.actionResponse.result', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_ACTIONRESPONSE_RESULT, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=185,
  serialized_end=327,
)

_ACTIONREQUEST.fields_by_name['corpus'].enum_type = _ACTIONREQUEST_CORPUS
_ACTIONREQUEST_CORPUS.containing_type = _ACTIONREQUEST
_ACTIONRESPONSE_RESULT.containing_type = _ACTIONRESPONSE
_ACTIONRESPONSE.fields_by_name['result'].message_type = _ACTIONRESPONSE_RESULT
DESCRIPTOR.message_types_by_name['actionRequest'] = _ACTIONREQUEST
DESCRIPTOR.message_types_by_name['actionResponse'] = _ACTIONRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

actionRequest = _reflection.GeneratedProtocolMessageType('actionRequest', (_message.Message,), {
  'DESCRIPTOR' : _ACTIONREQUEST,
  '__module__' : 'data_pb2'
  # @@protoc_insertion_point(class_scope:pb.actionRequest)
  })
_sym_db.RegisterMessage(actionRequest)

actionResponse = _reflection.GeneratedProtocolMessageType('actionResponse', (_message.Message,), {

  'Result' : _reflection.GeneratedProtocolMessageType('Result', (_message.Message,), {
    'DESCRIPTOR' : _ACTIONRESPONSE_RESULT,
    '__module__' : 'data_pb2'
    # @@protoc_insertion_point(class_scope:pb.actionResponse.Result)
    })
  ,
  'DESCRIPTOR' : _ACTIONRESPONSE,
  '__module__' : 'data_pb2'
  # @@protoc_insertion_point(class_scope:pb.actionResponse)
  })
_sym_db.RegisterMessage(actionResponse)
_sym_db.RegisterMessage(actionResponse.Result)



_FORMATDATA = _descriptor.ServiceDescriptor(
  name='FormatData',
  full_name='pb.FormatData',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=329,
  serialized_end=394,
  methods=[
  _descriptor.MethodDescriptor(
    name='DoFormat',
    full_name='pb.FormatData.DoFormat',
    index=0,
    containing_service=None,
    input_type=_ACTIONREQUEST,
    output_type=_ACTIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_FORMATDATA)

DESCRIPTOR.services_by_name['FormatData'] = _FORMATDATA

# @@protoc_insertion_point(module_scope)
