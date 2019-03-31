package component_test

var objectTemplateSimple string = `
{
	"Name": "Sample Object",
	"Components": [
		{"Type": "*go-component/component_test.FakeComponent"},
		{"Type": "*go-component/component_test.FakeConfiguredComponent", "Data": {
			"Items": [
				{"Id": "1", "Count": 1},
				{"Id": "2", "Count": 2},
				{"Id": "3", "Count": 3}
			]
		}}
	],
	"Objects": [
	]
}
`;

var objectTemplateNested string = `
{
	"Name": "Sample Object",
	"Components": [{
		"Type": "*go-component/component_test.FakeComponent"
	}],
	"Objects": [{
			"Name": "N/A",
			"Components": [{
				"Type": "*go-component/component_test.FakeComponent"
			}],
			"Objects": []
		}, {
			"Name": "One",
			"Objects": [{
				"Name": "Two",
				"Components": [{
					"Type": "*go-component/component_test.FakeComponent"
				}, {
					"Type": "*go-component/component_test.FakeConfiguredComponent",
					"Data": {
						"Items": [{
							"Id": "1",
							"Count": 1
						}, {
							"Id": "2",
							"Count": 2
						}, {
							"Id": "3",
							"Count": 3
						}]
					}
				}],
				"Objects": []
			}]
		}
	]
}
`;