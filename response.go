package main

var AboutInfo = []byte(`{
	"type": "carousel",
	"contents": [
	  {
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
	  },
	  {
		"type": "bubble",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "About Me",
			  "weight": "bold",
			  "size": "xxl",
			  "margin": "md",
			  "color": "#ffffffde"
			},
			{
			  "type": "separator",
			  "margin": "lg"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "Hi, my name is Peng-Jui Wang. I am currently a computer science student and an undergraduate researcher in MUI Lab.",
				  "size": "sm",
				  "color": "#ffffffde",
				  "flex": 0,
				  "wrap": true
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "My interests lie in the fields of HCI and web development. I create better systems that support human interaction with technology.",
				  "size": "sm",
				  "color": "#ffffffde",
				  "flex": 0,
				  "wrap": true
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "💭 Want to know more? ",
				  "size": "sm",
				  "color": "#ffffffde",
				  "flex": 0,
				  "wrap": true
				},
				{
				  "type": "text",
				  "text": "Click on the links below or interact with the bot!",
				  "size": "sm",
				  "color": "#ffffffde",
				  "flex": 0,
				  "wrap": true
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "margin": "xxl",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "button",
				  "style": "link",
				  "height": "sm",
				  "action": {
					"type": "uri",
					"label": "GitHub",
					"uri": "https://github.com/james5418"
				  },
				  "margin": "sm",
				  "color": "#ffffffde"
				},
				{
				  "type": "button",
				  "style": "link",
				  "height": "sm",
				  "color": "#ffffffde",
				  "action": {
					"type": "uri",
					"label": "LinkedIn",
					"uri": "http://www.linkedin.com/in/pjwang"
				  },
				  "margin": "sm"
				}
			  ]
			}
		  ],
		  "backgroundColor": "#464F69"
		},
		"styles": {
		  "footer": {
			"separator": true
		  }
		}
	  }
	]
  }`)

var UsageInfo = []byte(`{
	"type": "bubble",
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "Usage",
		  "weight": "bold",
		  "size": "xxl",
          "margin": "none",
		  "color": "#ffffffde"
		},
		{
		  "type": "separator",
		  "margin": "md"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "lg",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "The following commands are available.",
			  "size": "sm",
			  "color": "#ffffffde",
			  "wrap": true
			},
			{
			  "type": "text",
			  "text": "Use them to play with the bot !",
			  "size": "sm",
			  "color": "#ffffffde",
			  "wrap": true
			}
		  ]
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "margin": "none",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "style": "secondary",
			  "height": "sm",
			  "color": "#27ACB2",
			  "action": {
				"type": "message",
				"label": "about",
				"text": "about"
			  },
			  "margin": "lg"
			},
			{
			  "type": "button",
			  "style": "secondary",
			  "height": "sm",
			  "color": "#27ACB2",
			  "action": {
				"type": "message",
				"label": "skill",
				"text": "skill"
			  },
			  "margin": "lg"
			},
			{
			  "type": "button",
			  "style": "secondary",
			  "height": "sm",
			  "color": "#27ACB2",
			  "action": {
				"type": "message",
				"label": "project",
				"text": "project"
			  },
			  "margin": "lg"
			},
			{
			  "type": "button",
			  "style": "secondary",
			  "height": "sm",
			  "color": "#27ACB2",
			  "action": {
				"type": "message",
				"label": "experience",
				"text": "experience"
			  },
			  "margin": "lg"
			},
			{
			  "type": "button",
			  "style": "secondary",
			  "height": "sm",
			  "color": "#27ACB2",
			  "action": {
				"type": "message",
				"label": "help",
				"text": "help"
			  },
			  "margin": "lg"
			}
		  ]
		}
	  ],
	  "backgroundColor": "#464F69"
	},
	"styles": {
	  "footer": {
		"separator": true
	  }
	}
  }`)

var ExperinceInfo = []byte(`{
	"type": "carousel",
	"contents": [
	  {
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://static.wixstatic.com/media/421374_f6abec3fd1144c3894ce2ced5d30c563~mv2_d_7360_4912_s_4_2.jpg/v1/fill/w_862,h_600,fp_0.50_0.50,q_85,usm_0.66_1.00_0.01,enc_auto/421374_f6abec3fd1144c3894ce2ced5d30c563~mv2_d_7360_4912_s_4_2.jpg",
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
		  "spacing": "md",
		  "contents": [
			{
			  "type": "text",
			  "text": "Mobile and Ubiquitous Interaction (MUI) Lab",
			  "wrap": true,
			  "weight": "bold",
			  "gravity": "center",
			  "size": "xl",
			  "color": "#ffffffde"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "Researcher & Developer",
				  "size": "sm",
				  "color": "#999999",
				  "margin": "xs",
				  "flex": 0
				},
				{
				  "type": "text",
				  "text": "2021 ‑ Present  |  Hsinchu, Taiwan",
				  "size": "sm",
				  "color": "#999999",
				  "margin": "xs",
				  "flex": 0
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
				  "color": "#ffffffde",
				  "size": "md",
				  "flex": 1,
				  "text": "Working on mobile attention and notification, especially focus on notification sorting and summary.",
				  "wrap": true
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
				  "color": "#ffffffde",
				  "size": "md",
				  "flex": 1,
				  "text": "Also Developing notification management Apps at the same time.",
				  "wrap": true
				}
			  ]
			}
		  ],
		  "backgroundColor": "#464F69"
		}
	  },
	  {
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://png.pngtree.com/thumb_back/fw800/background/20190223/ourmid/pngtree-award-ceremony-black-gold-style-background-material-backgroundawards-partyawards-ceremonyend-image_71212.jpg",
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
		  "spacing": "md",
		  "contents": [
			{
			  "type": "text",
			  "text": "Honors & Awards",
			  "wrap": true,
			  "weight": "bold",
			  "gravity": "center",
			  "size": "xl",
			  "color": "#ffffffde"
			},
			{
			  "type": "separator",
			  "margin": "md"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "2022",
				  "size": "md",
				  "color": "#ffffffde",
				  "margin": "xs",
				  "decoration": "underline",
				  "style": "italic"
				},
				{
				  "type": "text",
				  "text": "Finalist, TSMC IT & Microsoft Careerhack",
				  "size": "md",
				  "color": "#ffffffde",
				  "margin": "xs",
				  "flex": 0,
				  "wrap": true
				}
			  ]
			},
			{
			  "type": "separator",
			  "margin": "xxl"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "md",
			  "contents": [
				{
				  "type": "text",
				  "text": "2021",
				  "size": "md",
				  "color": "#ffffffde",
				  "margin": "xs",
				  "style": "italic",
				  "decoration": "underline"
				},
				{
				  "type": "text",
				  "text": "Academic achievement award at NYCU",
				  "size": "md",
				  "color": "#ffffffde",
				  "margin": "xs",
				  "flex": 0,
				  "wrap": true
				}
			  ]
			}
		  ],
		  "backgroundColor": "#464F69"
		}
	  }
	]
  }`)

var SkillInfo = []byte(`{
	"type": "bubble",
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "md",
	  "action": {
		"type": "uri",
		"uri": "https://linecorp.com"
	  },
	  "contents": [
		{
		  "type": "text",
		  "text": "Skills",
		  "size": "xxl",
          "margin": "none",
		  "weight": "bold",
		  "color": "#ffffff"
		},
		{
		  "type": "separator",
		  "margin": "sm"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "icon",
				  "url": "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/twitter/322/speech-balloon_1f4ac.png"
				},
				{
				  "type": "text",
				  "text": "Languages",
				  "weight": "bold",
				  "margin": "sm",
				  "flex": 0,
				  "color": "#ffffffde"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "C, C++, Python, Javascript, GO",
				  "size": "sm",
				  "align": "start",
				  "color": "#ffffffde"
				}
			  ]
			}
		  ],
		  "margin": "lg"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "icon",
				  "url": "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/google/313/globe-with-meridians_1f310.png"
				},
				{
				  "type": "text",
				  "text": "Web Dev",
				  "weight": "bold",
				  "margin": "sm",
				  "flex": 0,
				  "color": "#ffffffde"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "Node.js, Express.js, Gin, Vue.js",
				  "size": "sm",
				  "align": "start",
				  "color": "#ffffffde"
				}
			  ]
			}
		  ],
		  "margin": "lg"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "icon",
				  "url": "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/microsoft/310/file-cabinet_1f5c4-fe0f.png"
				},
				{
				  "type": "text",
				  "text": "Database",
				  "weight": "bold",
				  "margin": "sm",
				  "flex": 0,
				  "color": "#ffffffde"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "MySQL, MongoDB, Redis",
				  "size": "sm",
				  "align": "start",
				  "color": "#ffffffde"
				}
			  ]
			}
		  ],
		  "margin": "lg"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "icon",
				  "url": "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/twitter/322/pick_26cf-fe0f.png"
				},
				{
				  "type": "text",
				  "text": "Tools",
				  "weight": "bold",
				  "margin": "sm",
				  "flex": 0,
				  "color": "#ffffffde"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "Git, Unix/Linux, Power BI, Android Studio",
				  "size": "sm",
				  "align": "start",
				  "color": "#ffffffde"
				}
			  ]
			}
		  ],
		  "margin": "lg"
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "icon",
				  "url": "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/twitter/322/magnifying-glass-tilted-left_1f50d.png"
				},
				{
				  "type": "text",
				  "text": "Research Interests",
				  "weight": "bold",
				  "margin": "sm",
				  "flex": 0,
				  "color": "#ffffffde"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "HCI, Mobile attention and notification",
				  "size": "sm",
				  "color": "#ffffffde"
				}
			  ]
			}
		  ],
		  "margin": "lg"
		}
	  ],
	  "backgroundColor": "#464F69"
	}
  }`)

var ProjectInfo = []byte(`{
	"type": "carousel",
	"contents": [
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://free-url-shortener.rb.gy/open-graph.png",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "RESTful API",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "URL Shortener",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Node.js / Express.js / Redis",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Offering URL shortening service"
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Applied unit test by Mocha and Chai"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			},
			{
			  "type": "button",
			  "style": "link",
			  "color": "#ffffffde",
			  "action": {
				"type": "uri",
				"label": "Read more ...",
				"uri": "https://github.com/james5418/URL_Shortener"
			  },
			  "margin": "sm",
			  "height": "sm"
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  },
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://github.com/scott306lr/fcs_orient_2021/blob/main/client/src/assets/map.jpg?raw=true",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "Web application",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "Orienteering Game",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Node.js / Express.js / MongoDB",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Website for NYCU orientation camp "
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Real‑time orientating game"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			},
			{
			  "type": "button",
			  "style": "link",
			  "color": "#ffffffde",
			  "action": {
				"type": "uri",
				"label": "Read more ...",
				"uri": "https://github.com/scott306lr/fcs_orient_2021"
			  },
			  "margin": "sm",
			  "height": "sm"
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  },
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://github.com/james5418/iWelfare/raw/main/frontend_vue/demo1/public/logo.png",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "Web application",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "iWelfare",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Vue.js / MySQL",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Website for searching social welfare"
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Allow users to quickly search for information through tags"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			},
			{
			  "type": "button",
			  "style": "link",
			  "color": "#ffffffde",
			  "action": {
				"type": "uri",
				"label": "Read more ...",
				"uri": "https://github.com/james5418/iWelfare"
			  },
			  "margin": "sm",
			  "height": "sm"
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  },
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://cdn.shopify.com/s/files/1/0312/5870/6055/products/RushH-5000-POG.jpg?v=1605821543",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "AI",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "Rush Hour Game Solver",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Python",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "AI solver for the classic board game ”Rush Hour”"
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Implemented by A* search algorithm"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			},
			{
			  "type": "button",
			  "style": "link",
			  "color": "#ffffffde",
			  "action": {
				"type": "uri",
				"label": "Read more ...",
				"uri": "https://github.com/james5418/Rush_Hour_solver"
			  },
			  "margin": "sm",
			  "height": "sm"
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  },
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://static.wixstatic.com/media/421374_3d46e662238d45eaa6bdaf547d588a86~mv2.jpg/v1/fill/w_481,h_417,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/iStock-861146256.jpg",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "text": "Android App",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "Noti-Manager",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Java / Android Studio",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "A more flexible interface to organize notifications"
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Automatically sorted and classified notifications by ML models"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  },
	  {
		"type": "bubble",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://www.tsmc.com/static/english/careers/2022Careerhack/images/A02.png",
				  "size": "full",
				  "aspectMode": "cover",
				  "aspectRatio": "150:100",
				  "gravity": "center",
				  "flex": 1
				},
				{
				  "type": "box",
				  "layout": "horizontal",
				  "contents": [
					{
					  "type": "text",
					  "size": "xs",
					  "color": "#ffffff",
					  "align": "start",
					  "wrap": true,
					  "text": "Dashboard"
					}
				  ],
				  "backgroundColor": "#EC3D44",
				  "paddingAll": "2px",
				  "paddingStart": "10px",
				  "paddingEnd": "10px",
				  "cornerRadius": "70px",
				  "offsetTop": "18px",
				  "offsetStart": "18px",
				  "flex": 0,
				  "position": "absolute"
				}
			  ]
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
					  "text": "Resource Recycling Platform",
					  "color": "#ffffff",
					  "weight": "bold"
					},
					{
					  "type": "text",
					  "text": "Microsoft Power BI",
					  "color": "#ffffffcc",
					  "size": "sm"
					}
				  ],
				  "spacing": "sm"
				},
				{
				  "type": "box",
				  "layout": "vertical",
				  "contents": [
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Visualized the cyclic process of resources"
					},
					{
					  "type": "text",
					  "contents": [],
					  "size": "sm",
					  "wrap": true,
					  "margin": "sm",
					  "color": "#ffffffde",
					  "text": "Offered a vehicle for users to share and claim resources"
					}
				  ],
				  "backgroundColor": "#ffffff1A",
				  "cornerRadius": "10px",
				  "margin": "xl",
				  "paddingAll": "10px"
				}
			  ]
			},
			{
			  "type": "button",
			  "style": "link",
			  "color": "#ffffffde",
			  "action": {
				"type": "uri",
				"label": "Read more ...",
				"uri": "https://www.tsmc.com/static/english/careers/2022Careerhack/main.html"
			  },
			  "margin": "sm",
			  "height": "sm"
			}
		  ],
		  "backgroundColor": "#464F69",
		  "paddingAll": "20px"
		}
	  }
	]
  }`)
