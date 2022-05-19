This is {{ .Adjective }} go template.

A favorite food of {{ read yaml "traits.yaml" "name" }} is {{ read yaml "traits.yaml" "favorite.food" }}.
A favorite number of {{ read yaml "traits.yaml" "name" }} is {{ read yaml "traits.yaml" "favorite.number" }}.
A favorite booleanValue of {{ read yaml "traits.yaml" "name" }} is {{ read yaml "traits.yaml" "favorite.booleanValue" }}.