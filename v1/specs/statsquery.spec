# Model
model:
  rest_name: statsquery
  resource_name: statsqueries
  entity_name: StatsQuery
  package: jenova
  description: |-
    StatsQuery is a generic API to retrieve time series data stored by the Aporeto
    system. The API allows different types of queries that are all protected within
    the namespace of the user.
  aliases:
  - sq
  create: true

# Attributes
attributes:
- name: results
  description: Results contains the result of the query.
  type: external
  exposed: true
  subtype: time_series_results
  read_only: true
  autogenerated: true
