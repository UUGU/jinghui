(function($){
    $.extend($,{
        imgLazyLoad:function(){
            var timer,
                len = $('img.lazyload').length;
            function getPos(node) {
                var scrollx = document.documentElement.scrollLeft || document.body.scrollLeft,
                    scrollt = document.documentElement.scrollTop || document.body.scrollTop;
                var pos = node.getBoundingClientRect();
                return {top:pos.top + scrollt, right:pos.right + scrollx, bottom:pos.bottom + scrollt, left:pos.left + scrollx }
            }
            function loading(){
                timer && clearTimeout(timer);
                timer = setTimeout(function(){
                    var scrollTop = document.documentElement.scrollTop || document.body.scrollTop,
                        imgs=$('img.lazyload');
                    screenHeight = document.documentElement.clientHeight;
                    for(var i = 0;i < imgs.length;i++){
                        var pos = getPos(imgs[i]),
                            posT = pos.top,
                            posB = pos.bottom,
                            screenTop = screenHeight+scrollTop;
                        if((posT > scrollTop && posT <  screenTop) || (posB > scrollTop && posB < screenTop)){
                            imgs[i].src = imgs[i].getAttribute('data-img');
                            $(imgs[i]).removeClass('lazyload');
                        }else{
                            // new Image().src = imgs[i].getAttribute('data-img');
                        }
                    }
                },100);
            }
            if(!len) return;
            loading();
            $(window).on('scroll resize',function(){
                if(!$('img.lazyload').length){
                    return;
                }else{
                    loading();
                }
            })
        }
    })
})($||Zepto||jQuery);


$(function() {

	// 区间筛选
	$("#screen").click(function(){
		var screen = $("#scrurl").val();
		var price_start = $("#price_start").val();
		var price_end = $("#price_end").val();
		var charges = $("#charges").val();
		var url = screen +'price_start=' + price_start + '&price_end='+ price_end + '&charges='+charges;
		location.href = url;

	});

	// 搜索区域
	$("#search").click(function(){
		$(':input[type="submit"]').prop('disabled', false);
		var kw = $("#kw").val();
		var search = document.filterForm.action;
		if(kw==""){
			var params = '?r=total';
			window.location.href = search + params;
			return false;	
		}
		var params = '?r=search/?kw='+ encodeURIComponent(kw) ;

		window.location.href = search + params;
		return false;
	});



	$("#tobuy").click(function(){
		order_link();

	});
	$("#img_tobuy").click(function(){
		order_link();

	});
	function order_link(){
		var gid  = $('#gid').val();
		var appid  = $('#appid').val();
		$.ajax({ 
			url: "http://capi.jingtuitui.com/jd/get_goods_link?", //用户处理页
			dataType: "jsonp",
			type: "POST",
			//传送请求数据
			data: 'gid='+ gid +'&appid='+ appid ,
			success: function(data) {
				window.location.href = data.result;
			},
		})
	}



	$(window).scroll(function () {
		if(getScrollTop() >= 100){
			var str = location.href;
			if(str.indexOf("search")!=-1){  
				$(".list_").css("top","51px");
			}else{
				$(".list_").css("top","0");
			}
			$(".list_").css("position","fixed");
			$(".list_").css("z-index","1");
			$(".list_").fadeIn(100); // 开始淡入
		} else{
			$(".list_").css("position","unset");
		}
		if(getScrollTop() > 500){
			$(".to_top").css("display","block");			
		}else{
			$(".to_top").css("display","none");	
		}
		
		
	});



	$(".to_top").click(function() {
		setScrollTop(0);
	});

	
/** 
 *获取scrollTop的值，兼容所有浏览器  
 */  
function getScrollTop() {  
    var scrollTop = document.documentElement.scrollTop || window.pageYOffset || document.body.scrollTop;  
    return scrollTop;  
}  
  
/** 
 *设置scrollTop的值，兼容所有浏览器  
 */  
function setScrollTop(scroll_top) {  
    document.documentElement.scrollTop = scroll_top;  
    window.pageYOffset = scroll_top;  
    document.body.scrollTop = scroll_top;  
} 








})
