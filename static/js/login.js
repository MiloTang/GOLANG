$().ready(function(){

$('#updatecode').click(function()
{
this.src="checkverifycode.php?"+Math.random();
});

$('#subform').click(function(){
var username=$('#username').val();
var password=$('#password').val();
var verifycode=$('#verifycode').val();
if (username.length<5) 
{
$("i:eq(0)").html("用户名不小于5位");
return false;
}
if (password.length<6) 
{
$("i:eq(1)").html("密码不小于6位");
return false;
}
if (verifycode.length<6) 
{
$("i:eq(2)").html("验证码不小于6位");
return false;
}
else
{
$.ajax
({ 
                    url: "Check.php", 
                    type: "GET",
					dataType: "json",
                    data:{"verifycode":verifycode},
					error: function()
					{  
						$("i:eq(2)").html("验证出现问题请稍候再试");
						return false;
					},  
                    success: function (result) 
					{ 
						if (result=="true") 
						{ 
							$('#signupForm').submit();
							return true;
						}
						else
						{
						   $("i:eq(2)").html("验证码不正确");
						   return false;
						}	
                    }
                });
}

});

$('#username').focus(function ()
{
$("i:eq(0)").html("");
});
$('#password').focus(function ()
{
$("i:eq(1)").html("");
});
$('#verifycode').focus(function ()
{
$("i:eq(2)").html("");
});
});
