$(function(){
    $(".si_serviceleftLast").hover(function(){
    $(".si_serviceleftLast div").fadeIn(0);
    },function(){
        $(".si_serviceleftLast div").fadeOut(0);
    })
    $(".si_serviceleftLI2").hover(function(){
    $(".si_serviceleftLI2 div").fadeIn(0);
    },function(){
        $(".si_serviceleftLI2 div").fadeOut(0);
    })
	
	  $(".si_serviceleftLI1").hover(function(){
    $(".si_serviceleftLI1 div").fadeIn(0);
    },function(){
        $(".si_serviceleftLI1 div").fadeOut(0);
    })
    $("#gototop").click(function(){ 
        $("html,body").animate({scrollTop :0}, 300); 
        return false; 
    })
    $("#gotofoot").click(function(){ 
        $("html,body").animate({scrollTop : $(document).height()}, 300); 
        return false; 
    })
	$(".si_serviceleft2 ul li").hover(function(){
        $(this).css('background-color','Transparent');
        },function(){
        $(this).css('background-color','');
        })
	})