<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 14px;">组织架构管理</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div class="pull-left" style="height: 44px; line-height: 44px; width: 260px;">
        <span style="font-size: 10px;font-weight: 600;height: 30px; line-height: 30px; margin-top: 7px;display: inline"
              class="pull-left">&nbsp;&nbsp;所属域：</span>
        <select id="h-org-domain-list" class="form-control pull-left"
                style="width: 180px;height: 24px; line-height: 24px; margin-top: 10px;padding: 0px;">
        </select>
    </div>
    <div class="pull-right">
        <button onclick="OrgObj.add()" class="btn btn-info btn-sm" title="新增机构信息">
            <i class="icon-plus"> 新增</i>
        </button>
        <button onclick="OrgObj.upload()" class="btn btn-info btn-sm" title="导入机构信息">
        <span class="icon-edit"> 导入</span>
        </button>
        <button onclick="OrgObj.download()" class="btn btn-info btn-sm" title="导出机构信息">
        <span class="icon-trash"> 导出</span>
        </button>
        <button onclick="OrgObj.edit()" class="btn btn-info btn-sm" title="编辑机构信息">
            <span class="icon-edit"> 编辑</span>
        </button>
        <button onclick="OrgObj.delete()" class="btn btn-danger btn-sm" title="删除机构信息">
            <span class="icon-trash"> 删除</span>
        </button>
    </div>
</div>
<div class="row" style="padding-top: 6px;">
    <div class="col-sm-12 col-md-5 col-lg-3">
        <div id="h-org-tree-info" class="thumbnail">
            <div class="col-ms-12 col-md-12 col-lg-12">
                <div style="border-bottom: #598f56 solid 1px;height: 44px; line-height: 44px;">
                    <div class="pull-left">
                        <span><i class="icon-sitemap"> </i>组织架构树</span>
                    </div>
                    <div class="pull-right">
                        <span>
                            <i class=" icon-search" style="margin-top: 15px;"></i>&nbsp;
                    </span>
                    </div>
                </div>
            </div>
            <div id="h-org-tree-info-list" class="col-sm-12 col-md-12 col-lg-12"
                 style="padding:15px 5px;overflow: auto">

            </div>
        </div>
    </div>
    <div id="h-org-table-info" class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <table id="h-org-info-table-details"
               data-toggle="table"
               data-striped="true"
               data-unique-id="org_id"
               data-url="/api/orgunits"
               data-side-pagination="client"
               data-pagination="true"
               data-page-size="20"
               data-page-list="[20, 50, 100, 200]"
               data-search="false">
            <thead>
            <tr>
                <th data-field="state" data-checkbox="true"></th>
                <th data-field="id">机构编码</th>
                <th data-field="name">机构描述</th>
                <th data-field="up_org_id" data-formatter="OrgObj.upOrgId">上级机构编码</th>
                <th data-align="center" data-field="create_date">创建日期</th>
                <th data-align="center" data-field="create_user">创建人</th>
                <th data-align="center" data-field="modify_date">修改日期</th>
                <th data-align="center" data-field="modify_user">修改人</th>
            </tr>
            </thead>
        </table>
    </div>
</div>
<script type="text/javascript">

    var OrgObj = {
        $table:$('#h-org-info-table-details'),
        /*
        * 新增机构信息,只能在自己拥有写入权限的域中新增机构信息
        * */
        add:function(){
            $.Hmodal({
                body:$("#org_input_form").html(),
                height:"300px",
                header:"新增机构",
                callback:function(hmode){
                    $.HAjaxRequest({
                        url:"/api/orgunits",
                        data:$("#h-org-add-info").serialize(),
                        type:"post",
                        success:function (data) {
                            $.Notify({
                                title:"操作成功",
                                message:"插入机构信息成功",
                                type:"success",
                            });
                            $(hmode).remove();
                            var domain_id = $("#h-org-domain-list").val()
                            OrgObj.tree(domain_id)
                        },
                    })
                },
                preprocess:function(){
                    console.log("hi")
            
                    var domain_id = $("#h-org-domain-up-id").val()
                    $.getJSON("/api/orgunits",{domain_id:domain_id},function(data){
                        var arr = new Array()
                        $(data).each(function(index,element){
                            var ijs = {}
                            ijs.id=element.id;
                            ijs.text=element.name;
                            ijs.upId=element.parent_id;
                            arr.push(ijs)
                        });

                        var ijs = {};
                        ijs.id="root_vertex_system";
                        ijs.text="机构树根节点";
                        ijs.upId="######hzwy23#####";
                        arr.push(ijs)

                        $("#h-org-up-id").Hselect({
                            data:arr,
                            height:"30px",
                            value:$("#h-org-tree-info-list").attr("data-selected"),
                        })
                    })

                }
            })
        },
        /*
        * 编辑处理函数,
        * 当右侧table中没有机构被选中时,默认会编辑左侧被选中的机构
        * 如果左侧也没有机构被选中,则提示没有任何被选中的机构
        * 这个函数只在edit函数中被调用
        * */
        handle_edit : function (row) {
            var domain_id = row[0].domain_id;
            $.Hmodal({
                body:$("#org_modify_form").html(),
                header:"修改机构信息",
                height:"360px",
                preprocess:function () {
                    /*
                     * 初始化下拉框中机构信息
                     * */
                    $.getJSON("/api/orgunits",{domain_id:domain_id},function(data){
                        var arr = new Array()
                        $(data).each(function(index,element){
                            var ijs = {}
                            ijs.id=element.id;
                            ijs.text=element.name;
                            ijs.upId=element.parent_id;
                            arr.push(ijs)
                        });

                        var ijs = {};
                        ijs.id="root_vertex_system";
                        ijs.text="机构树根节点";
                        ijs.upId="######hzwy23#####";
                        arr.push(ijs)

                        $("#h-modify-org-up-id").Hselect({
                            data:arr,
                            value:row[0].parent_id,
                            height:"30px",
                        });
                    });

                    /*
                     * 在编辑框中，填上目前的机构信息。
                     * */
                    var org_id = row[0].id

                    var org_name = row[0].name

                    $("#h-modify-org-id").val(org_id)
                    $("#h-modify-org-name").val(org_name)

                },
                callback:function(hmode){
                    $.HAjaxRequest({
                        url:"/api/orgunits/update",
                        type:"put",
                        data:$("#h-org-modify-info").serialize(),
                        success:function (data) {
                            $.Notify({
                                title:"温馨提示：",
                                message:"修改机构信息成功",
                                type:"success",
                            })
                            $(hmode).remove();
                            OrgObj.tree(domain_id)
                        },
                    })
                }
            })
        },
        /*
        * 机构编辑按钮,当点击页面上的就编辑按钮时,
        * 会调用此函数
        * */
        edit:function(){
            var row = OrgObj.$table.bootstrapTable("getSelections").concat();

            if (row.length == 0){
                var selected_id = $("#h-org-tree-info-list").attr("data-selected");
                if (selected_id == undefined){
                    $.Notify({
                        title:"温馨提示",
                        message:"请在列表中选择一个需要编辑的机构",
                        type:"warning",
                    });
                    return
                }
                row.push(OrgObj.$table.bootstrapTable('getRowByUniqueId',selected_id))
                OrgObj.handle_edit(row);
                return
            } else if (row.length == 1){
                OrgObj.handle_edit(row);
            } else {
                $.Notify({
                    title:"温馨提示",
                    message:"您在列表中选中了多个机构，不知道需要编辑哪一个",
                    type:"warning",
                })
            }
        },
        /*
        * 在table中获取某个些机构的全部子机构信息
        * */
        getSubOrg:function(set,all){

            var addArray = new Array();

            var search = function(obj,arr){
                $(arr).each(function(index,element){
                    if (obj.org_id == element.up_org_id){
                        addArray.push(element)
                        var newArray = all.splice(index,1)
                        search(element,newArray)
                    }
                });
            };

            $(set).each(function(index,element){
                search(element,all)
            });

            return addArray
        },
        delete:function(){
            var domain_id = $("#h-org-domain-list").val();

            var data = OrgObj.$table.bootstrapTable("getSelections").concat();

            if (data.length == 0) {
                var selected_id = $("#h-org-tree-info-list").attr("data-selected");
                if (selected_id == undefined) {
                    $.Notify({
                        title: "温馨提示",
                        message: "请在列表中选择一个需要编辑的机构",
                        type: "warning",
                    });
                    return
                }
                data.push(OrgObj.$table.bootstrapTable('getRowByUniqueId', selected_id))
            }

            $.Hconfirm({
                callback:function () {
                    $.HAjaxRequest({
                        url:"/api/orgunits/delete",
                        type:"post",
                        data:{JSON:JSON.stringify(data),domain_id:domain_id},
                        success:function () {
                            $.Notify({
                                title:"操作成功",
                                message:"删除机构信息成功",
                                type:"success",
                            });
                            OrgObj.tree(data[0].domain_id)
                        },
                    })
                },
                body:"确认要删除选中的机构信息吗"
            })
        },
        download:function(){
            var domain_id = $("#h-org-domain-list").val()
            var x=new XMLHttpRequest();
            x.open("GET", "/api/orgunits/download?domain_id="+domain_id, true);
            x.responseType = 'blob';
            x.onload=function(e){
                download(x.response, "机构信息.xlsx", "application/vnd.ms-excel" );
            }
            x.send();
        },
        upload:function(param){
            $.Hupload({
                url:"/api/orgunits/upload",
                header:"导入机构信息",
                callback:function () {
                    var domain_id = $("#h-org-domain-list").val();
                    OrgObj.tree(domain_id)
                },
            })
        },
        tree:function(domain_id){
          $.getJSON("/api/orgunits",{domain_id:domain_id},function(data){
              if (data.length==0){
                  $.Notify({
                      title:"温馨提示",
                      message:"您选择的域中没有机构信息",
                      type:"info",
                  });
                  OrgObj.$table.bootstrapTable('load',[])
                  $("#h-org-tree-info-list").Htree({
                      data:[],
                  })
              } else {
                  var arr = new Array()
                  $(data).each(function(index,element){
                      var ijs = {}
                      ijs.id = element.id
                      ijs.text = element.name
                      ijs.upId = element.parent_id
                      arr.push(ijs)
                  });
                  $("#h-org-tree-info-list").Htree({
                      data:arr,
                      onChange:function(obj){
                          var id = $(obj).attr("data-id")
                          $.getJSON("/api/orgunits",{parent_id:id},function(data){
                              OrgObj.$table.bootstrapTable('load',data)
                          });
                      }
                  });
                  OrgObj.$table.bootstrapTable('load',data)
              }
          })
        },
        upOrgId:function(value,row,index){
            var upcombine = row.parent_id
            if (upcombine.length==2){
                return upcombine[1]
            }else{
                return upcombine
            }
        }
    };

    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        $("#h-org-tree-info").height(hwindow-139);
        $("#h-org-table-info").height(hwindow-139);
        $("#h-org-tree-info-list").height(hwindow-213);

        // $.getJSON("/v1/auth/domain/owner",function(data){
        //     var arr = new Array()
        //     $(data).each(function(index,element){
        //         var ijs = {}
        //         ijs.id=element.domain_id
        //         ijs.text=element.domain_desc
        //         ijs.upId="####hzwy23###"
        //         arr.push(ijs)
        //     });
        //     $("#h-org-domain-list").Hselect({
        //         data:arr,
        //         height:"24px",
        //         width:"180px",
        //         onclick:function () {
        //             var id = $("#h-org-domain-list").val();
        //             OrgObj.tree(id);
        //         },
        //     });
        //     $.getJSON("/api/orgunits",function (data) {
        //         OrgObj.tree(data);
        //         $("#h-org-domain-list").val(data).trigger("change")
        //         $('#h-org-info-table-details').bootstrapTable({
        //             height:hwindow-130,
        //             queryParams:function (params) {
        //                 params.domain_id = $("#h-org-domain-list").val();
        //                 return params
        //             },
        //         });
        //     });

        // });
        OrgObj.tree({});
        $('#h-org-info-table-details').bootstrapTable({
            height:hwindow-130,
            queryParams:function (params) {
                params.domain_id = $("#h-org-domain-list").val();
                return params
            },
        });
    });
</script>

<script type="text/html" id="org_input_form">
    <form class="row"  id="h-org-add-info">
        <div class="col-sm-12 col-md-12 col-lg-12">
            <div class="form-group-sm col-sm-6 col-md-6 col-lg-6">
                <label class="h-label" style="width: 100%;">组织部门名称：</label>
                <input placeholder="请输入1-60位汉字，字母，数字（必填）" type="text" class="form-control" name="name" style="width: 100%;height: 30px;line-height: 30px;">
            </div>
        </div>
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 15px;">
            <div class="form-group-sm col-sm-6 col-md-6 col-lg-6">
                <label class="h-label" style="width: 100%;">上级组织部门：</label>
                <select id="h-org-up-id" name="parent_id" style="width: 100%;height: 30px;line-height: 30px;">
                </select>
            </div>
        </div>
    </form>
</script>

<script type="text/html" id="org_modify_form">
    <form id="h-org-modify-info">
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 18px;">
            <label class="h-label" style="width: 100%;">组织部门名称：</label>
            <input id="h-modify-org-name" placeholder="user name" type="text" class="form-control" name="name" style="width: 100%;height: 30px;line-height: 30px;">
        </div>
        <div class="col-sm-12 col-md-12 col-lg-12" style="margin-top: 18px;">
            <label class="h-label" style="width: 100%;">上级组织部门代码：</label>
            <select id="h-modify-org-up-id" name="parent_id" style="width: 100%;height: 30px;line-height: 30px;">
            </select>
        </div>
    </form>
</script>