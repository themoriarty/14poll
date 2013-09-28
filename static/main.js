$(document).ready(main);

function main(){
    $(".vote div").click(function(){
	if (!$(this).hasClass("disabled")){
	    $(this).toggleClass("disabled");
	    var vote = {poll: $(this).closest(".poll").attr("data-name"),
			option: $(this).closest(".vote").attr("data-name"),
			choice: $(this).attr("data-vote")};
	    console.log();
	}
    });
}