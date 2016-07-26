    
    ;(function($){

        $.ajax({

            url:"127.0.0.1:10000/podlist",  
            type:"GET",  
            dataType:"json",
            success:function(data){   
                var othml="";
                $.each(data.yeepay,function(i,v){        
                        othml+="<tr><td><a href='#'>"+v.appHealthz.podsAvailable+"</a></td><td class='tables-tds'><p><span>"+v.appName+"</span></p><p><span>"+v.appLabel.appversion+"</span></p></td><td>"+v.appDatacenter+"</td><td>"+v.appReplicas+"</td><td>"+v.appWorktime+"</td><td><span>更多</span></td></tr>"});
                $("#uls").html(othml);
                
            },
            error:function(){
                alert("请求失败");
            }
        })
                
                
            
       
    })(jQuery);