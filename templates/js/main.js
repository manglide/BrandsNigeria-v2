$(document).ready(function() {
	//$(".dropdown-button").dropdown({hover: false});
	$(function () {
    const ele = document.getElementById('ipl-progress-indicator')
    if (ele) {
      setTimeout(() => {
        ele.classList.add('available')
        setTimeout(() => {
        ele.outerHTML = ''
      }, 2000)
      }, 1000)
    }
	 });
	$("#datacomments").submit(function(event){
            event.preventDefault();
            var latitude = 0
			var longitude = 0
			if(navigator.geolocation){
	            	navigator.geolocation.getCurrentPosition(function(position){
	                	latitude = position.coords.latitude
					longitude = position.coords.longitude
					// do what you like with the input
					$inputLat = $('<input type="hidden" type="text" name="latitude"/>').val(latitude);
					$inputLon = $('<input type="hidden" type="text" name="longitude"/>').val(longitude);
					// append to the form
					$("#datacomments").append($inputLat);
					$("#datacomments").append($inputLon);
		            $.ajax({
		                    url:'/u/comments',
		                    type:'POST',
		                    data:$("#datacomments").serialize(),
		                    before: function() {
			                    	
		                    },
		                    success:function(result){
								console.log(result)
		                        // location.reload();
		                    },
		                    error:function(result){
		                    	   if(result.status == 401 && result.statusText === "Unauthorized") {
									alert("Sorry, you must be logged in to comment")
							   } else {
									alert(result.responseJSON.message)
							   } 
		                    }
		            });
	            	});		
			} else {
					$inputLat = $('<input type="hidden" type="text" name="latitude"/>').val(latitude);
					$inputLon = $('<input type="hidden" type="text" name="longitude"/>').val(longitude);
					// append to the form
					$("#datacomments").append($inputLat);
					$("#datacomments").append($inputLon);
		            $.ajax({
		                    url:'/u/comments',
		                    type:'POST',
		                    data:$("#datacomments").serialize(),
		                    before: function() {
			                    	
		                    },
		                    success:function(result){
								console.log(result)
		                        // location.reload();
		                    },
		                    error:function(result){
		                    	   if(result.status == 401 && result.statusText === "Unauthorized") {
									alert("Sorry, you must be logged in to comment")
							   } else {
									alert(result.responseJSON.message)
							   } 
		                    }
		            });
			}
			
     });
	
	 $('span.rates').each(function(i, obj) {
		 var rateVal = $(this).attr("data");
		 var rateValMod = rateVal % 1;
		 if(rateValMod !== 0) {
			 if($(this).parent().attr("id") !== undefined) {
				 var pT = document.getElementById($(this).parent().attr("id"));
				 $(pT).append("<span class='fa fa-star-half-full rates' />");
			 } 
		 }
	 });
	 $('span.ratesPage').each(function(i, obj) {
		 var rateVal = $(this).attr("data");
		 var rateValMod = rateVal % 1;
		 if(rateValMod !== 0) {
			 if($(this).parent().attr("id") !== undefined) {
				 var pT = document.getElementById($(this).parent().attr("id"));
				 $(pT).append("<i style='font-size:38px;color:#000;' class='fa fa-star-half-full' />");
			 }
		 }
	 });

	 $('span.commentrates').each(function(i, obj) {
		 var rateVal = $(this).attr("data");
		 var rateValMod = rateVal % 1;
		 if(rateValMod !== 0) {
			 if($(this).parent().attr("id") !== undefined) {
				 var pT = document.getElementById($(this).parent().attr("id"));
				 $(pT).append("<i style='font-size:18px;color:#000;' class='fa fa-star-half-full' />");
			 }
		 }
	 });
	
	$("#sendC").click(function(event){
            event.preventDefault();
            
            var comm = document.getElementById("comment")
            comm = $(comm).val();
            var aut = $("input[name=authorx]").val()
            var s = document.getElementById("rating")
            s = $(s).val()
            var str = $("#datacomments").serialize()
            var t = str.split("&")
            var sentiment = $("input[name='sentiment']:checked").val();
            
            var x = t[0] + "&" + t[1] + "&author=" + aut + "&sentiment="+ sentiment + "&comment=" + comm + "&rating=" + s
            var latitude = 0
			var longitude = 0
			if(comm == "") {
				alert("Please set your comment")
			} else if (aut == "") {
				alert("Please set the author value")
			} else if (sentiment == undefined) {
				alert("Please choose sentiment")
			} else {
				if(navigator.geolocation){
	            	navigator.geolocation.getCurrentPosition(function(position){
	                	latitude = position.coords.latitude
					longitude = position.coords.longitude
					// do what you like with the input
					$inputLat = $('<input type="hidden" type="text" name="latitude"/>').val(latitude);
					$inputLon = $('<input type="hidden" type="text" name="longitude"/>').val(longitude);
					// append to the form
					$("#datacomments").append($inputLat);
					$("#datacomments").append($inputLon);
		            $.ajax({
		                    url:'/u/comments',
		                    type:'POST',
		                    data:x,
		                    before: function() {
			                    	$("#sendC").text('Sending...please wait').addClass("disabled")
		                    },
		                    success:function(result){
								alert("Review submitted successfully")
								$("#sendC").text('Submit').addClass("disabled")
								location.reload();
		                    },
		                    error:function(result){
		                    	   if(result.status == 401 && result.statusText === "Unauthorized") {
									alert("Sorry, you must be logged in to comment")
							   } else {
									alert(result.responseJSON.message)
							   } 
		                    }
		            });
	            	});		
			} else {
					$inputLat = $('<input type="hidden" type="text" name="latitude"/>').val(latitude);
					$inputLon = $('<input type="hidden" type="text" name="longitude"/>').val(longitude);
					// append to the form
					$("#datacomments").append($inputLat);
					$("#datacomments").append($inputLon);
		            $.ajax({
		                    url:'/u/comments',
		                    type:'POST',
		                    data:x,
		                    before: function() {
			                    	$("#sendC").text('Sending...please wait').addClass("disabled")
		                    },
		                    success:function(result){
								alert("Review submitted successfully")
								$("#sendC").text('Submit').addClass("disabled")
								location.reload();
		                    },
		                    error:function(result){
		                    	   if(result.status == 401 && result.statusText === "Unauthorized") {
									alert("Sorry, you must be logged in to comment")
							   } else {
									alert(result.responseJSON.message)
							   } 
		                    }
		            });
			}	
			}
			
     });

	 // Get Data For Like Chart
	 var productsreviewlikes = document.getElementById("productsreviewlikes");
	 var products = $(productsreviewlikes).attr("data").split(/[,]+/).join();
	 makeBarChartHighLikes(products);
	 // Get Data For DisLike Chart
	 var productsreviewdislikes = document.getElementById("productsreviewdislikes");
	 var productsDislikes = $(productsreviewdislikes).attr("data").split(/[,]+/).join();
	 makeBarChartHighDisLikes(productsDislikes);
	 // Get Data For DisLike Chart
	 var productsreviewrating = document.getElementById("productsreviewrating");
	 var productsRating = $(productsreviewrating).attr("data").split(/[,]+/).join();
	 makeBarChartHighRating(productsRating);
	 // Get Data For Competitor 1
	 var firstCompetitor = document.getElementById("firstCompetitor");
	 var comp1data = $(firstCompetitor).attr("data");
	 comp1data = comp1data.trim();
	 if (/\s/.test(comp1data)) {
    	var cleanURL = comp1data.split(/\s/);
			cleanURL = cleanURL.join("-");
			cleanURL = cleanURL.toLocaleLowerCase();
			loadCompetitors1(cleanURL);
	} else {
			comp1data = comp1data.toLocaleLowerCase();
			loadCompetitors1(comp1data);
	}
	 // Get Data For Competitor 1
	 var secondCompetitor = document.getElementById("secondCompetitor");
	 var comp2data = $(secondCompetitor).attr("data");
	 comp2data = comp2data.trim();
	 if (/\s/.test(comp2data)) {
    	var cleanURL2 = comp2data.split(/\s/);
			cleanURL2 = cleanURL2.join("-");
			cleanURL2 = cleanURL2.toLocaleLowerCase();
			loadCompetitors1(cleanURL2);
	 } else {
			comp2data = comp2data.toLocaleLowerCase();
			loadCompetitors1(comp2data);
	 }
	 // Get Data Product ID for Map from Element AreaAcceptanceMapElem
	 var dataAcceptance = document.getElementById("AreaAcceptanceMapElem");
	 var accData = $(dataAcceptance).attr("data");
	 preMapCallAccept(accData);
	 // Get Data Product ID for Map from Element AreaRejectionMapElem
	 var dataRejection = document.getElementById("AreaRejectionMapElem");
	 var rejData = $(dataRejection).attr("data");
	 preMapCallReject(rejData);
	 // Get Recommendation
	 var productsrecommendation = document.getElementById("productsrecommendation");
	 var prRec = $(productsrecommendation).attr("data");
	 prRecommended(prRec);
});
