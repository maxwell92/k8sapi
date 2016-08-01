$(document).ready(function(){
  // ajax 分页 + 渲染数据  
    var num = 5;//每页显示的个数
    var n = 0;
    var m = -num;
    function ajax(pageType){
        var oul = $(".box").find("ul");
        var ohtml = "";
        $.ajax({
            type:"get",
          //  url:"data.json",
            url:"./applist",
            dataType:"json",
            success:function(data){
                $(oul).empty();
                if(n < data.length && pageType=="next"){ //上一页
                    n += num;
                    m += num;
                }else if(m > 0 && pageType=="prev"){ //下一页
                    n -= num;
                    m -= num;
                }else if(pageType=="first"){ //第一页
                    n = num;
                    m = 0;
                }else if(pageType=="last"){ //最后一页
                    n = data.length+(data.length%num)-1;
                    m = data.length+(data.length%num)-6;
                }
                $.each(data,function(i,val){
                    
                        ohtml += "<tr><td><a href='#'>"+val.Healthz.PodsAvailable+"</a></td><td class='tables-tds'><p><span>"+val.Name+"</span></p><p><span>"+val.Label+"</span></p></td><td>"+val.Datacenter+"</td><td>"+val.Replicas+"</td><td>"+val.Worktime+"</td><td><span>更多</span></td></tr>";
                    
                });
                $(".box").html(ohtml);
            }
        });
    };
    $(".next").click(function(){
        ajax("next");
    });
    $(".prev").click(function(){
        ajax("prev");
    });
    $(".first").click(function(){
        ajax("first");
    });
    $(".last").click(function(){
        ajax("last");
    });
    $(function(){ //初始化
        ajax("next");
    });
    


    
    //  默认Dashboard样式
    
  // $(".yp-leftbox-uls").find("li:first-child").css({"background":"#4842b7","color":"#fff"});
   

    //  滑过侧栏样式
   /* $(".yp-leftbox-uls li").hover(function(){
        $(this).css({"background":"#4842b7","color":"#fff"}).siblings().css({"background":"#fff","color":"#000"})
    })*/

    $(".yp-leftbox-uls li a").toggle(function(){
         
        $(this).siblings().show(300);

    },function(){
        $(this).siblings().hide(300);
    })












//  button 点击不冒泡
$("button").click(function(){
    return false;
})


//  复选框内容显示
$(".post-right dl dt input").toggle(function(){
    console.log(12345)
    $(this).parent().next().hide();
},function(){
    $(this).parent().next().show();
})



//   数据中心卡片
$(".post-r-namber p input").click(function(){
   console.log(234567)
    if($(this).attr("checked")=="checked"){
        console.log(123)
        $(this).next().addClass("add-sj-tap");
    }else{
        $(this).next().removeClass("add-sj-tap")
    }
})


$(".post-r-format p input").click(function(){
//   console.log(234567)
    if($(this).attr("checked")){
     //   $('input:radio[name="post-radio"]:checked').val()
        console.log(123)
        $(this).parent().addClass("post-format").siblings().removeClass("post-format");
    }
})


$(".post-btn1").click(function(){
    console.log(234);
    console.log($(this).parent().length);
    $(this).parent().append("<dd class='post-append'><input type='text' />：<input type='text' /><b class='post-btn2'>-</b></dd>")
})

$(".post-btn2").live("click",function(){
    
    $(this).parent().remove();
})







//    -- 上传


function doUpload() {  
     var formData = new FormData($( "#uploadForm" )[0]);  
     $.ajax({  
          url: 'http://www.dropzonejs.com/#with-component' ,  
          type: 'POST',  
          data: formData,  
          async: false,  
          cache: false,  
          contentType: false,  
          processData: false,  
          success: function (returndata) {  
              alert("上传成功！");  
          },  
          error: function (returndata) {  
              alert('错误数据，上传地址出现问题！');  
          } 
     });  
}
function change(){
    var a = $(".filename").val();
    console.log(a);
    if(a == "请选择文件..."){
        alert("请选择文件！");
        return false
    }else{
        doUpload()
        return true;
    }
}



$(function(){
    $("input[type=file]").change(function(){
        $(this).parents(".uploader").find(".filename").val($(this).val());
    });
    
    $("input[type=file]").each(function(){
        if($(this).val()==""){
            $(this).parents(".uploader").find(".filename").val("请选择文件...");
        }
    });
});





//   -- end






//    -- 高级选项


$(".post-super").toggle(function(){
    $(".post-r-pans").show();
},function(){
    $(".post-r-pans").hide();
})




$("")














































    
   
});
