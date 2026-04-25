---
title: TIDE API
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# TIDE API

Base URLs:

# Authentication

- HTTP Authentication, scheme: bearer

# Default

## GET Location

GET /api/location

Retrieves Location data from name or GeoJSON point.

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|name|query|string| no |Filters by location's name|
|point|query|string| no |Filters by locations geographical position (100km)|

> Response Examples

> 200 Response

```json
[
  {
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
    "marineId": "25",
    "name": "string",
    "point": "POINT(23.33 -11.2)",
    "meanSeaLevel": 1.28,
    "timezone": "string"
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|*anonymous*|[[Location](#schemalocation)]|false|none||none|
|» id|string(uuid)|true|none||Location ID|
|» marineId|string|true|none||ID from Marinha do Brasil|
|» name|string|true|none||Location name|
|» point|string|true|none||Geographical point [POINT(longitude latitude)]|
|» meanSeaLevel|number(float)|true|none||Mean Sea Level (m)|
|» timezone|string|true|none||Location's timezone|

## GET Tide

GET /api/location/{location.id}/tides/{day}

Retrieve Tide data from a specific Location and date

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|location.id|path|string(uuid)| yes |Location ID|
|day|path|string(date)| yes |Date for locate Tide Data|

> Response Examples

> 200 Response

```json
[
  {
    "time": "2019-08-24T14:15:22Z",
    "height": -9.99,
    "type": "HIGH"
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|*anonymous*|[[Tide](#schematide)]|false|none||none|
|» time|string(date-time)|true|none||Tide time|
|» height|number(float)|true|none||Tide height in meters (m)|
|» type|string|true|none||HIGH if height >= Location.meanSeaLevel, LOW otherwise|

#### Enum

|Name|Value|
|---|---|
|type|HIGH|
|type|LOW|

# Data Schema

<h2 id="tocS_Location">Location</h2>

<a id="schemalocation"></a>
<a id="schema_Location"></a>
<a id="tocSlocation"></a>
<a id="tocslocation"></a>

```json
{
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
  "marineId": "25",
  "name": "string",
  "point": "POINT(23.33 -11.2)",
  "meanSeaLevel": 1.28,
  "timezone": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|string(uuid)|true|none||Location ID|
|marineId|string|true|none||ID from Marinha do Brasil|
|name|string|true|none||Location name|
|point|string|true|none||Geographical point [POINT(longitude latitude)]|
|meanSeaLevel|number(float)|true|none||Mean Sea Level (m)|
|timezone|string|true|none||Location's timezone|

<h2 id="tocS_Tide">Tide</h2>

<a id="schematide"></a>
<a id="schema_Tide"></a>
<a id="tocStide"></a>
<a id="tocstide"></a>

```json
{
  "time": "2019-08-24T14:15:22Z",
  "height": -9.99,
  "type": "HIGH"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|time|string(date-time)|true|none||Tide time|
|height|number(float)|true|none||Tide height in meters (m)|
|type|string|true|none||HIGH if height >= Location.meanSeaLevel, LOW otherwise|

#### Enum

|Name|Value|
|---|---|
|type|HIGH|
|type|LOW|

