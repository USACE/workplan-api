package static

// Offices is raw JSON for offices
var Offices = []byte(`[
  {
    "id": "0088df5f-ec58-4654-9b71-3590266b475c",
    "symbol": "MVS",
    "name": "St. Louis District"
  },
  {
    "id": "0360bb56-92f8-4c1e-9b08-8396b216f2d3",
    "symbol": "LRDO",
    "name": "Ohio River Region"
  },
  {
    "id": "07c1c91e-2e42-4fbc-a550-f775a27eb419",
    "symbol": "NWK",
    "name": "Kansas City District"
  },
  {
    "id": "098c898f-7f0f-44e8-b2c1-329d7dd50166",
    "symbol": "SAW",
    "name": "Wilmington District"
  },
  {
    "id": "0cac2b45-a5df-49b3-9176-7d5145681958",
    "symbol": "NAD",
    "name": "North Atlantic Division"
  },
  {
    "id": "1f579664-d1db-4ee9-897e-47c16dc55012",
    "symbol": "NWO",
    "name": "Omaha District"
  },
  {
    "id": "2222f2f5-d512-41ee-83d7-3a6cfcbf5bfb",
    "symbol": "SPD",
    "name": "South Pacific Division"
  },
  {
    "id": "26dab361-76a0-4cb2-b5a5-01667ab7f7da",
    "symbol": "MVN",
    "name": "New Orleans District"
  },
  {
    "id": "26e6c300-480b-4e22-afae-7cc27dd9b116",
    "symbol": "SWT",
    "name": "Tulsa District"
  },
  {
    "id": "33f03e9a-711b-41e7-9bdd-66152b69128d",
    "symbol": "MVP",
    "name": "St. Paul District"
  },
  {
    "id": "4142c26c-0407-41ad-b660-8657ddb2be69",
    "symbol": "SAJ",
    "name": "Jacksonville District"
  },
  {
    "id": "4f4f6899-cac1-402f-adee-8109d5bc5db3",
    "symbol": "SWG",
    "name": "Galveston District"
  },
  {
    "id": "4ffaa895-0f05-4b59-8d12-86c901e2f229",
    "symbol": "LRN",
    "name": "Nashville District"
  },
  {
    "id": "586ac79a-083e-4c8c-8438-9585a88a4b3d",
    "symbol": "LRE",
    "name": "Detroit District"
  },
  {
    "id": "60754640-fef3-429b-b2f4-efdcf3b61e55",
    "symbol": "SWL",
    "name": "Little Rock District"
  },
  {
    "id": "60c46088-9c98-4927-991d-0bd126bbb62e",
    "symbol": "LRB",
    "name": "Buffalo District"
  },
  {
    "id": "64fb2c2f-e59a-44cc-a54d-22e8d7c909a0",
    "symbol": "NAE",
    "name": "New England District"
  },
  {
    "id": "6e3c2e48-a15a-4892-ba91-caab86499abc",
    "symbol": "NAP",
    "name": "Philadelphia District"
  },
  {
    "id": "74f7b6ff-b026-4272-b44a-cb1d536e1b8d",
    "symbol": "POD",
    "name": "Pacific Ocean Division"
  },
  {
    "id": "76bb611e-3fbf-4779-b251-203de3502670",
    "symbol": "SPN",
    "name": "San Francisco District"
  },
  {
    "id": "788cc853-235a-4e43-a136-b1f190a6a656",
    "symbol": "LRH",
    "name": "Huntington District"
  },
  {
    "id": "790ec8cf-8dad-48c9-bea9-9b8c26d29471",
    "symbol": "SAD",
    "name": "South Atlantic Division"
  },
  {
    "id": "7d7e962d-e554-48f0-8f82-b762a31441a6",
    "symbol": "NAB",
    "name": "Baltimore District"
  },
  {
    "id": "7fd53614-f484-4dfd-8fc4-d11b11fb071c",
    "symbol": "NAO",
    "name": "Norfolk District"
  },
  {
    "id": "834ea54d-2454-425f-b443-50ba4ab46e28",
    "symbol": "NWW",
    "name": "Walla Walla District"
  },
  {
    "id": "85ba21d8-ba4b-4060-a519-a3e69c1e29ed",
    "symbol": "NWDP",
    "name": "Pacific Northwest Region"
  },
  {
    "id": "89a1fe0c-03f3-47cf-8ee1-cd3de2e1ba7b",
    "symbol": "SPL",
    "name": "Los Angeles District"
  },
  {
    "id": "8fc88b15-9cd4-4e86-8b8c-6d956926010b",
    "symbol": "MVM",
    "name": "Memphis District"
  },
  {
    "id": "90173658-2de9-4329-926d-176c1b29089a",
    "symbol": "NWDM",
    "name": "Missouri River Region"
  },
  {
    "id": "90b958ea-0076-4925-87d8-670eb7da5551",
    "symbol": "SAS",
    "name": "Savannah District"
  },
  {
    "id": "99322682-b22f-4c47-972a-81c4d782b0d5",
    "symbol": "SAC",
    "name": "Charleston District"
  },
  {
    "id": "99a6b349-535a-4aab-b742-9bdd145461e7",
    "symbol": "LRDG",
    "name": "Great Lakes Region"
  },
  {
    "id": "9a631b0c-d8ad-4411-8220-04683c9c24f4",
    "symbol": "POH",
    "name": "Hawaii District"
  },
  {
    "id": "a0baec43-2817-4161-b654-c3c513b5276b",
    "symbol": "POA",
    "name": "Alaska District"
  },
  {
    "id": "a222e733-2fa7-4cd8-b3a6-065956e693f0",
    "symbol": "SWF",
    "name": "Fort Worth District"
  },
  {
    "id": "a9929dc4-7d7c-4ddb-b727-d752137ffc10",
    "symbol": "SPA",
    "name": "Albuquerque District"
  },
  {
    "id": "b952664c-4b11-4d85-89fa-a2cc405b1131",
    "symbol": "LRL",
    "name": "Louisville District"
  },
  {
    "id": "b9c56905-9dad-4418-9654-d1fcd9b3a57f",
    "symbol": "MVR",
    "name": "Rock Island District"
  },
  {
    "id": "c18588b6-25ab-42c9-b31a-a33c585d0b49",
    "symbol": "NWP",
    "name": "Portland District"
  },
  {
    "id": "c88758e9-4575-44b0-9d38-b6c0ee909061",
    "symbol": "NWS",
    "name": "Seattle District"
  },
  {
    "id": "d02f876f-eb00-425b-aeca-09fa105d5bc2",
    "symbol": "MVK",
    "name": "Vicksburg District"
  },
  {
    "id": "d0b7ddca-a321-44bd-bf2c-059c9c8cbe23",
    "symbol": "LRD",
    "name": "Great Lakes and Ohio River Division"
  },
  {
    "id": "d3da00c9-f839-4add-90a9-73053292d196",
    "symbol": "NWD",
    "name": "Northwestern Division"
  },
  {
    "id": "dd580032-c210-4f98-8ab7-bda92ff2fe5e",
    "symbol": "MVD",
    "name": "Mississippi Valley Division"
  },
  {
    "id": "df64a2de-91a2-4e6c-85c9-53d3d03f6794",
    "symbol": "SPK",
    "name": "Sacramento District"
  },
  {
    "id": "e7c9cfe8-99eb-4845-a058-46e53a75b28b",
    "symbol": "LRP",
    "name": "Pittsburgh District"
  },
  {
    "id": "eb545b18-5498-43c8-8652-f73e16446cc0",
    "symbol": "SAM",
    "name": "Mobile District"
  },
  {
    "id": "ede616b6-5ab7-42c6-9489-7c09bfb6a54b",
    "symbol": "NAN",
    "name": "New York District"
  },
  {
    "id": "fa9b344c-911c-43e9-966c-b0e1357e385c",
    "symbol": "LRC",
    "name": "Chicago District"
  },
  {
    "id": "fe551ee7-3b04-440c-89a4-162dffd99ed2",
    "symbol": "SWD",
    "name": "Southwestern Division"
  },
  {
    "id": "cf0e2fde-9156-4a96-a8bc-32640cb0043d",
    "symbol": "STUDY",
    "name": "Columbia  Study"
  },
  {
    "id": "e303450f-af6a-4272-a262-fdb94f8e3e86",
    "symbol": "SERFC",
    "name": "Southeast River Forecast Center"
  }
]`);