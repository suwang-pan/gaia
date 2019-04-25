# Model
model:
  rest_name: tag
  resource_name: tags
  entity_name: Tag
  package: tagle
  group: core/tag
  description: |-
    A tag is a string in the form of "key=value" that can applied to all objects in
    the system. They are used for policy resolution. Tags starting by a "$" are
    derived from the property of an object (for instance an object with ID set to
    xxx and a name set to "the name" will be tagged by default with "$name=the name"
    and "$id=xxx"). Tags starting with an "@" have been generated by an external
    system.
  extends:
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: count
    description: Count represents the number of time the tag is used.
    type: integer
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: namespace
    description: Namespace represents the namespace of the counted tag.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    primary_key: true

  - name: value
    description: Value represents the value of the tag.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_chars: ^[\w\d\*\$\+\.:,|@<>/-]+=[= \-\/\!\?\{\}\(\)\w\d\*\$\+\.:;,|@%&~<>#/"]+$
    allowed_chars_message: must contain at least one '=' symbol separating two valid
      words.
    example_value: key=value
    primary_key: true
    validations:
    - $tag
