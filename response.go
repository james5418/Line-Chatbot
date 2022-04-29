package main

var UsageInfo = []byte(`{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "Help Message",
		  "weight": "bold",
		  "size": "xl"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Place",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Time",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "10:00 - 23:00",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "sm",
	  "contents": [
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "CALL",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "WEBSITE",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [],
		  "margin": "sm"
		}
	  ],
	  "flex": 0
	}
  }`)

var ExperinceInfo = []byte(`{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "Experience Message",
		  "weight": "bold",
		  "size": "xl"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Place",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Time",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "10:00 - 23:00",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "sm",
	  "contents": [
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "CALL",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "WEBSITE",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [],
		  "margin": "sm"
		}
	  ],
	  "flex": 0
	}
  }`)

var AboutInfo = []byte(`{
	"type": "bubble",
	"header": {
	  "type": "box",
	  "layout": "horizontal",
	  "contents": [
		{
		  "type": "image",
		  "url": "https://media-exp1.licdn.com/dms/image/C5603AQHsFj8YCdlCQg/profile-displayphoto-shrink_800_800/0/1648265862269?e=1656547200&v=beta&t=JQdFlWcZsWo8qROr-yLL_jJhRa7M2MQic1nBtl6Xo-o",
		  "size": "full",
		  "aspectMode": "cover",
		  "aspectRatio": "140:120",
		  "gravity": "center",
		  "flex": 1
		}
	  ],
	  "paddingAll": "0px"
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "text",
				  "contents": [],
				  "size": "xl",
				  "wrap": true,
				  "text": "Peng-Jui, Wang",
				  "color": "#ffffff",
				  "weight": "bold"
				},
				{
				  "type": "text",
				  "text": "Junior in Computer Science at NYCU",
				  "color": "#ffffffcc",
				  "size": "sm",
				  "contents": []
				},
				{
				  "type": "text",
				  "contents": [],
				  "size": "sm",
				  "wrap": true,
				  "margin": "lg",
				  "color": "#ffffffde",
				  "text": "Backend Developer 🗄️"
				},
				{
				  "type": "text",
				  "contents": [],
				  "size": "sm",
				  "wrap": true,
				  "margin": "lg",
				  "color": "#ffffffde",
				  "text": "HCI Researcher 🧑"
				}
			  ],
			  "spacing": "sm"
			}
		  ]
		}
	  ],
	  "paddingAll": "20px",
	  "backgroundColor": "#464F69"
	}
  }`)

var SkillInfo = []byte(`{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "Skill Message",
		  "weight": "bold",
		  "size": "xl"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Place",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Time",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "10:00 - 23:00",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "sm",
	  "contents": [
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "CALL",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "WEBSITE",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [],
		  "margin": "sm"
		}
	  ],
	  "flex": 0
	}
  }`)

var ProjectInfo = []byte(`{
	"type": "bubble",
	"hero": {
	  "type": "image",
	  "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"uri": "http://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "Project Message",
		  "weight": "bold",
		  "size": "xl"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Place",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Time",
				  "color": "#aaaaaa",
				  "size": "sm",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "10:00 - 23:00",
				  "wrap": true,
				  "color": "#666666",
				  "size": "sm",
				  "flex": 5
				}
			  ]
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "sm",
	  "contents": [
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "CALL",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "button",
		  "style": "link",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "WEBSITE",
			"uri": "https://linecorp.com"
		  }
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [],
		  "margin": "sm"
		}
	  ],
	  "flex": 0
	}
  }`)
