# Model
model:
  rest_name: connectionsresult
  resource_name: connectionsresults
  entity_name: ConnectionsResult
  package: guy
  group: core/enforcer
  description: A result from a connections request.
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: connections
    description: Contains a batch of connections.
    type: refList
    exposed: true
    subtype: currentconnection
    extensions:
      refMode: pointer

  - name: requestID
    description: Unique ID generated for each request.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: totalConnections
    description: Contains the total number of connections for the connection request.
    type: integer
    exposed: true
    omit_empty: true
