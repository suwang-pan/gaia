{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "represents data received from the service",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "data",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "POST_HOOK",
                "PRE_HOOK"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Defines the type of the hook",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "type",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "hks",
            "hk"
        ],
        "create": null,
        "delete": false,
        "description": "Hook to integrate in an Aporeto service",
        "entity_name": "Hook",
        "extends": [
            "@described",
            "@disabled",
            "@named"
        ],
        "get": false,
        "package": null,
        "resource_name": "hooks",
        "rest_name": "hook",
        "root": null,
        "update": false
    }
}