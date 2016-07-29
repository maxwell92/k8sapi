    
    ;(function($){

        $.ajax({

            url:"http://localhost:10000/getapplist",
            type:"GET",  
            dataType:"json",
            success:function(data){   
                var othml="";
                $.each(data,function(i,v){        
                        othml+="<tr><td><a href='#'>"+v.Healthz.PodsAvailable+"</a></td><td class='tables-tds'><p><span>"+v.Name+"</span></p><p><span>"+v.Label+"</span></p></td><td>"+v.Datacenter+"</td><td>"+v.Replicas+"</td><td>"+v.Worktime+"</td><td><span>更多</span></td></tr>"});
                $("#uls").html(othml);
                
            },
            error:function(){
                alert("请求失败");
            }
        })
                
                
            
       
    })(jQuery);
