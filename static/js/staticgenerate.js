$().ready(function()
{
$('#homelist').click(function(){
$.ajax
({ 
                    url: "GenerateHtml.php", 
                    type: "GET",
					dataType: "json",
                    data:{"generatetype":"homelist"},
					error: function()
					{  
						alert("失败");
						return false;
					},  
                    success: function (result) 
					{ 
						if (result=="failed") 
						{ 
							alert(result);
							return false;
						}
						else
						{
						  $("i:eq(0)").html(result);
						   return true;
						}	
                    }
                });
});

$('#newartcontentlist').click(function(){
$.ajax
({ 
                    url: "GenerateHtml.php", 
                    type: "GET",
					dataType: "json",
                    data:{"generatetype":"newartcontentlist"},
					error: function()
					{  
						alert("失败");
						return false;
					},  
                    success: function (result) 
					{ 
						if (result=="failed") 
						{ 
							alert(result);
							return false;
						}
						else
						{
						   $("i:eq(0)").html(result);
						   return true;
						}	
                    }
                });
});

$('#alldailycontentlist').click(function(){

$.ajax
({ 
                    url: "GenerateHtml.php", 
                    type: "GET",
					dataType: "json",
                    data:{"generatetype":"alldailycontentlist"},
					error: function()
					{  
						alert("失败");
						return false;
					},  
                    success: function (result) 
					{ 
						if (result=="failed") 
						{ 
							alert(result);
							return false;
						}
						else
						{
						  $("i:eq(0)").html(result);
						   return true;
						}	
                    }
                });
});

$('#other').click(function(){

$.ajax
({ 
                    url: "GenerateHtml.php", 
                    type: "GET",
					dataType: "json",
                    data:{"generatetype":"other"},
					error: function()
					{  
						alert("失败");
						return false;
					},  
                    success: function (result) 
					{ 
						if (result=="failed") 
						{ 
							alert(result);
							return false;
						}
						else
						{
						  alert(result);
						   return true;
						}	
                    }
                });
});

});
