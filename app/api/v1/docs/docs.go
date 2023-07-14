// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.hitosea.com/",
        "contact": {},
        "license": {
            "name": "AGPL-3.0 license",
            "url": "http://www.gnu.org/licenses/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/doo/project/list": {
            "get": {
                "description": "获取项目列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Dootask"
                ],
                "summary": "获取项目列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/doo/user/info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Dootask"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/doo/user/list": {
            "get": {
                "description": "获取用户列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Dootask"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/analyze/complete": {
            "get": {
                "description": "OKR整体平均完成度",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OkrAnalyze"
                ],
                "summary": "OKR整体平均完成度",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/interfaces.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/interfaces.OkrAnalyzeOverall"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/okr/analyze/dept/complete": {
            "get": {
                "description": "OKR各部门平均完成度",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OkrAnalyze"
                ],
                "summary": "OKR各部门平均完成度",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/interfaces.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/interfaces.OkrAnalyzeDept"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/okr/analyze/dept/score": {
            "get": {
                "description": "OKR各部门评分分布",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OkrAnalyze"
                ],
                "summary": "OKR各部门评分分布",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/interfaces.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/interfaces.OkrAnalyzeScoreDept"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/okr/analyze/personnel": {
            "get": {
                "description": "OKR人员评分率",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OkrAnalyze"
                ],
                "summary": "OKR人员评分率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/interfaces.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/interfaces.OkrAnalyzePersonnel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/okr/analyze/score": {
            "get": {
                "description": "OKR评分分布",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OkrAnalyze"
                ],
                "summary": "OKR评分分布",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/interfaces.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/interfaces.OkrAnalyzeScore"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/okr/confidence/update": {
            "post": {
                "description": "更新信心指数",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "更新信心指数",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "信心指数",
                        "name": "confidence",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/create": {
            "post": {
                "description": "创建OKR",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "创建OKR",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.OkrCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/department/list": {
            "get": {
                "description": "部门OKR列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "部门OKR列表",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.OkrDepartmentListReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/finish": {
            "post": {
                "description": "结束/重启目标",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "结束/重启目标",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "状态 0-重启 1-结束",
                        "name": "finished",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/follow": {
            "get": {
                "description": "关注或取消关注目标",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "关注或取消关注目标",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "目标id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/follow/list": {
            "get": {
                "description": "关注的OKR列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "关注的OKR列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "目标",
                        "name": "objective",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/my/list": {
            "get": {
                "description": "我的OKR列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "我的OKR列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "目标",
                        "name": "objective",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/participant/list": {
            "get": {
                "description": "参与的OKR列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "参与的OKR列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "目标",
                        "name": "objective",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/participant/update": {
            "post": {
                "description": "更新参与人",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "更新参与人",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "参与人,多个用逗号隔开",
                        "name": "participant",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/replay/create": {
            "post": {
                "description": "添加复盘",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "添加复盘",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.OkrReplayCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/replay/detail": {
            "get": {
                "description": "复盘详情",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "复盘详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "复盘id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/replay/list": {
            "get": {
                "description": "OKR复盘列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "OKR复盘列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "目标",
                        "name": "objective",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/score": {
            "post": {
                "description": "OKR评分",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "OKR评分",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "个人评分",
                        "name": "score",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/update": {
            "post": {
                "description": "更新OKR",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "更新OKR",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/interfaces.OkrUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        },
        "/okr/update/progress": {
            "post": {
                "description": "更新进度和进度状态",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Okr"
                ],
                "summary": "更新进度和进度状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "进度",
                        "name": "progress",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "进度状态 0-默认 1-正常 2-有风险 3-延期 4-已结束",
                        "name": "status",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/interfaces.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "interfaces.OkrAnalyzeDept": {
            "type": "object",
            "properties": {
                "complete": {
                    "description": "okr 完成数",
                    "type": "integer"
                },
                "department_id": {
                    "description": "部门id",
                    "type": "integer"
                },
                "department_name": {
                    "description": "部门名称",
                    "type": "string"
                },
                "total": {
                    "description": "okr 总数",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrAnalyzeOverall": {
            "type": "object",
            "properties": {
                "complete": {
                    "description": "okr 完成数",
                    "type": "integer"
                },
                "total": {
                    "description": "okr 总数",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrAnalyzePersonnel": {
            "type": "object",
            "properties": {
                "complete": {
                    "description": "okr 完成数",
                    "type": "integer"
                },
                "total": {
                    "description": "okr 总数",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrAnalyzeScore": {
            "type": "object",
            "properties": {
                "seven_to_ten": {
                    "description": "7-10分",
                    "type": "integer"
                },
                "three_to_seven": {
                    "description": "3-7分",
                    "type": "integer"
                },
                "total": {
                    "description": "okr 总数",
                    "type": "integer"
                },
                "unscored": {
                    "description": "未评分",
                    "type": "integer"
                },
                "zero_to_three": {
                    "description": "0-3分",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrAnalyzeScoreDept": {
            "type": "object",
            "properties": {
                "department_id": {
                    "description": "部门id",
                    "type": "integer"
                },
                "department_name": {
                    "description": "部门名称",
                    "type": "string"
                },
                "seven_to_ten": {
                    "description": "7-10分",
                    "type": "integer"
                },
                "three_to_seven": {
                    "description": "3-7分",
                    "type": "integer"
                },
                "total": {
                    "description": "okr 总数",
                    "type": "integer"
                },
                "unscored": {
                    "description": "未评分",
                    "type": "integer"
                },
                "zero_to_three": {
                    "description": "0-3分",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrCreateReq": {
            "type": "object",
            "required": [
                "align_objective",
                "ascription",
                "end_at",
                "key_results",
                "priority",
                "project_id",
                "start_at",
                "title",
                "type",
                "visible_range"
            ],
            "properties": {
                "align_objective": {
                    "description": "对齐目标",
                    "type": "string"
                },
                "ascription": {
                    "description": "归属 1-部门 2-个人",
                    "type": "integer"
                },
                "end_at": {
                    "description": "结束时间",
                    "type": "string"
                },
                "key_results": {
                    "description": "关键结果",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/interfaces.OkrKeyResultCreateReq"
                    }
                },
                "priority": {
                    "description": "优先级",
                    "type": "string"
                },
                "project_id": {
                    "description": "项目id",
                    "type": "integer"
                },
                "start_at": {
                    "description": "开始时间",
                    "type": "string"
                },
                "title": {
                    "description": "目标",
                    "type": "string"
                },
                "type": {
                    "description": "类型 1-承诺型 2-挑战型",
                    "type": "integer"
                },
                "visible_range": {
                    "description": "可见范围  1-全公司 2-仅相关成员 3-仅部门成员",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrDepartmentListReq": {
            "type": "object",
            "properties": {
                "completed": {
                    "description": "是否已完成未评分 0-未完成 1-已完成",
                    "type": "integer"
                },
                "department_id": {
                    "description": "部门id",
                    "type": "integer"
                },
                "end_at": {
                    "description": "结束时间",
                    "type": "string"
                },
                "objective": {
                    "description": "目标",
                    "type": "string"
                },
                "start_at": {
                    "description": "开始时间",
                    "type": "string"
                },
                "type": {
                    "description": "类型 1-承诺型 2-挑战型",
                    "type": "integer"
                },
                "userid": {
                    "description": "用户id",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrKeyResultCreateReq": {
            "type": "object",
            "required": [
                "confidence",
                "end_at",
                "participant",
                "start_at",
                "title"
            ],
            "properties": {
                "confidence": {
                    "description": "信心指数",
                    "type": "integer"
                },
                "end_at": {
                    "description": "结束时间",
                    "type": "string"
                },
                "participant": {
                    "description": "参与人,多个用逗号隔开",
                    "type": "string"
                },
                "start_at": {
                    "description": "开始时间",
                    "type": "string"
                },
                "title": {
                    "description": "关键结果",
                    "type": "string"
                }
            }
        },
        "interfaces.OkrKeyResultUpdateReq": {
            "type": "object",
            "required": [
                "confidence",
                "end_at",
                "id",
                "participant",
                "start_at",
                "title"
            ],
            "properties": {
                "confidence": {
                    "description": "信心指数",
                    "type": "integer"
                },
                "end_at": {
                    "description": "结束时间",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "participant": {
                    "description": "参与人,多个用逗号隔开",
                    "type": "string"
                },
                "start_at": {
                    "description": "开始时间",
                    "type": "string"
                },
                "title": {
                    "description": "关键结果",
                    "type": "string"
                }
            }
        },
        "interfaces.OkrReplayComment": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "comment": {
                    "description": "评价",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                }
            }
        },
        "interfaces.OkrReplayCreateReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "comments": {
                    "description": "复盘评价",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/interfaces.OkrReplayComment"
                    }
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "problem": {
                    "description": "问题与不足",
                    "type": "string"
                },
                "value": {
                    "description": "价值与收获",
                    "type": "string"
                }
            }
        },
        "interfaces.OkrUpdateReq": {
            "type": "object",
            "required": [
                "align_objective",
                "end_at",
                "id",
                "key_results",
                "priority",
                "project_id",
                "start_at",
                "title",
                "type",
                "visible_range"
            ],
            "properties": {
                "align_objective": {
                    "description": "对齐目标",
                    "type": "string"
                },
                "end_at": {
                    "description": "结束时间",
                    "type": "string"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "key_results": {
                    "description": "关键结果",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/interfaces.OkrKeyResultUpdateReq"
                    }
                },
                "priority": {
                    "description": "优先级",
                    "type": "string"
                },
                "project_id": {
                    "description": "项目id",
                    "type": "integer"
                },
                "start_at": {
                    "description": "开始时间",
                    "type": "string"
                },
                "title": {
                    "description": "目标",
                    "type": "string"
                },
                "type": {
                    "description": "类型 1-承诺型 2-挑战型",
                    "type": "integer"
                },
                "visible_range": {
                    "description": "可见范围  1-全公司 2-仅相关成员 3-仅部门成员",
                    "type": "integer"
                }
            }
        },
        "interfaces.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态, [200=成功, 400=失败, 401=未登录, 403=无相关权限, 404=请求接口不存在, 405=请求方法不允许, 500=系统错误]",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "msg": {
                    "description": "信息",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "http://localhost",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "DooTask",
	Description: "DooTask是一款轻量级的开源在线项目任务管理工具，提供各类文档协作工具、在线思维导图、在线流程图、项目管理、任务分发、即时IM，文件管理等工具。",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
