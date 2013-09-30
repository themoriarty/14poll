$(document).ready(main);

function voteUrl(poll){
    return "/poll/" + poll + "/vote"
}

function main(){
    $(".vote div").click(function(){
	if (!$(this).hasClass("disabled")){
	    var self = $(this);
	    $(this).toggleClass("disabled");
	    var poll_id = $(this).closest(".poll").attr("data-name");
	    $.ajax({
		type: "POST",
		url: voteUrl(poll_id),
		data: {csrf_token: g_token,
		       option: $(this).closest(".vote").attr("data-name"),
		       choice: $(this).attr("data-vote")},
		success: function(data){
		    console.log(data);
		    self.closest(".vote").slideUp();
		},
		error: function(r, status){
		    console.error(status);
		    self.toggleClass("disabled");
		}
	    });
	}
    });
}