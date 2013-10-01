$(document).ready(main);

function voteUrl(poll){
    return "/poll/" + poll + "/vote"
}

function checkToShowResults(poll){
    if (poll.find(".vote:visible[data-name]").length == 0){
	return true;
    }
    return false;
}

function getOption(btn){
    if (btn.closest(".option").find("input").length == 1){
	return [btn.closest(".option").find("input").val(), true];
    }
    return [btn.closest(".vote").attr("data-name"), false];
}

function sortResults(){
    var indexes = new Array();
    $(".results").each(function(i){
	indexes.push(i);
    });
    function getScore(i, kind){
	return $($(".results")[i]).find("." + kind).text() * 1;
    }
    function cmp(i1, i2, kind){
	var v1 = getScore(i1, kind);
	var v2 = getScore(i2, kind);
	if (v1 < v2){
	    return -1;
	} else if (v1 == v2){
	    return 0;
	}
	return 1;
    }
    indexes.sort(function (i1, i2){
	var ret = cmp(i1, i2, "VoteAgainst");
	if (ret == 0){
	    ret = -cmp(i1, i2, "VoteFor");
	}
	return ret;
    });
    indexes.forEach(function(e, i){
	indexes[i] = $($(".option")[e]).clone();
    });
    indexes.forEach(function(e, i){
	$($(".option")[i]).replaceWith(e);
    });
}

function main(){
    $(".vote .mybtn").click(function(){
	if (!$(this).hasClass("disabled")){
	    var self = $(this);
	    $(this).toggleClass("disabled");
	    var poll_id = $(this).closest(".poll").attr("data-name");
	    var option = getOption(self);
	    $.ajax({
		type: "POST",
		url: voteUrl(poll_id),
		data: {csrf_token: g_token,
		       option: option[0],
		       choice: $(this).attr("data-vote")},
		success: function(data){
		    console.log(data);
		    self.closest(".vote").slideUp(400, function(){
			if (option[1] || checkToShowResults(self.closest(".poll"))){
			    document.location.reload(); // XXX
			}
		    });
		},
		error: function(r, status){
		    console.error(status);
		    self.toggleClass("disabled");
		}
	    });
	}
    });
    var linkRegex = /[a-z]+:\/\/\S+/i; // this has a common problem with trailing "." and ","
    $(".name").each(function(idx, e){
	// this replaces only the first URL
	var text = $(e).text();
	var url = text.match(linkRegex);
	if (url){
	    var parts = text.split(linkRegex, 1);
	    $(e).text(parts[0]);
	    var a = document.createElement("a");
	    $(a).attr("href", url).attr("target", "_blank");
	    $(a).text(url);
	    aaa = a;
	    $(e).append(a);
	    if (parts.length > 1){
		$(e).append(document.createTextNode(parts[1]));
	    }
	}
    });
    if ($(".vote").length < 2){
	sortResults();
    }
}