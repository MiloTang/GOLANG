$(function(){
	// ??????
	$("#floatPanel a.arrow").eq(0).click(function(){
		$("html,body").animate({scrollTop :0}, 300);
		return false;
	});
	$("#floatPanel a.arrow").eq(1).click(function(){
		$("html,body").animate({scrollTop : $(document).height()}, 300);
		return false;
	});

	var panel1 = $(".popPanel1");	
	var w = panel1.outerWidth();
	
	$(".qrcode").hover(function(){
		panel1.css("width","0px").show();
		panel1.animate({"width" : w + "px"},300);
	},function(){
		panel1.animate({"width" : "0px"},300);
	});
	var panel2 = $(".popPanel2");	
	var w = panel2.outerWidth();
	
	$(".qrcode1").hover(function(){
		panel2.css("width","0px").show();
		panel2.animate({"width" : w + "px"},300);
	},function(){
		panel2.animate({"width" : "0px"},300);
	});
	
});