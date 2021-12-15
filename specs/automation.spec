# Model
model:
  rest_name: automation
  resource_name: automations
  entity_name: Automation
  package: sephiroth
  group: integration/automation
  description: |-
    Allows you to define some JavaScript code and specify the conditions under which
    it should
    be executed.
  aliases:
  - autos
  - auto
  get:
    description: Retrieves the automation with the given ID.
  update:
    description: Updates the automation with the given ID.
  delete:
    description: Deletes the automation with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'
  - '@signed'
  validations:
  - $automations

# Attributes
attributes:
  v1:
  - name: actions
    description: Contains the code that will be executed if the condition is met.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: aporetoToken
    description: |-
      Contains the Microsegmentation token used by the automation's HTTP client. This
      token is derived from the automation's app credential attribute.
    type: string
    stored: true
    autogenerated: true
    encrypted: true

  - name: appCredential
    description: |-
      Contains the app credential associated with the automation which has its roles
      deduced from the automation's entitlements.
    type: string
    stored: true
    autogenerated: true
    encrypted: true

  - name: condition
    description: |-
      Condition contains the code that will be executed to decide if any action(s)
      should be executed. Providing a condition for an automation with a
      "Webhook" trigger type will have no impact as the condition will not be
      evaluated. If no condition is defined, then the automation action(s) will be
      executed; this behaves akin to a condition that always succeeds.
    type: string
    exposed: true
    stored: true
    example_value: 'function when(m, params) { return { continue: true }}'

  - name: entitlements
    description: Declares which operations are allowed on which identities.
    type: external
    exposed: true
    subtype: _automation_entitlements
    stored: true

  - name: errors
    description: Contains the error of the last run.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    autogenerated: true

  - name: events
    description: |-
      Contains the identity and operation an event must have to trigger the
      automation.
    type: external
    exposed: true
    subtype: _automation_events
    stored: true

  - name: immediateExecution
    description: |-
      If set and the trigger is of type Time, the automation will be run at create or
      update before being scheduled.
    type: boolean
    exposed: true
    stored: true

  - name: lastExecTime
    description: The last successful execution tine.
    type: time
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: parameters
    description: Contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true

  - name: schedule
    description: |-
      Specifies when to run the automation. Must be in valid CRON format. This
      only applies if the trigger is set to `Time`.
    type: string
    exposed: true
    stored: true

  - name: stdout
    description: Contains the standard output of the last run.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: token
    description: |-
      Holds the unique access token used as a password to trigger the
      authentication. It will be visible only after creation.
    type: string
    exposed: true
    stored: true
    autogenerated: true
    transient: true

  - name: tokenRenew
    description: If set to `true` a new token will be issued and the previous one invalidated.
    type: boolean
    exposed: true

  - name: trigger
    description: Controls when the automation should be triggered.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Event
    - RemoteCall
    - Webhook
    - Time
    default_value: Time
    orderable: true

# Relations
relations:
- rest_name: trigger
  get:
    description: |-
      Allows a system to trigger the automation if its `trigger` property is set
      to `RemoteCall`.
  create:
    description: |-
      Allows a system to trigger the automation if its `trigger` property is set
      to `RemoteCall`.
