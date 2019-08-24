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
	
	function openShare(url, urlToPost, width, height) {
			window.open(url, '_blank', 'width=' + width + ',height=' + height + ',scrollbars=no,status=no');
			return false;
	}
	
	$(".socialb").click(function(){
		var link = $(this).attr("data-href")
		window.open('fb-messenger://share?link=' + encodeURIComponent(link) + '&app_id=491555231663023');
	})
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
								alert("Review submitted successfully")
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
		                    data:$("#datacomments").serialize(),
		                    before: function() {
			                    	
		                    },
		                    success:function(result){
								alert("Review submitted successfully")
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
	$("#storeblogimg").click(function(event){
		$(this).empty().html('Uploading...please wait')
		var data = new FormData();
	    $.each($('#imageblog')[0].files, function(i, file) {
	        data.append('file-'+i, file);
	    });
		
		$.ajax({
	        url: '/sendBlogImageForUpload',
			type: 'POST',
	        data: data,
	        cache: false,
	        contentType: false,
	        processData: false,
			before: function() {
			    $("#storeblogimg").empty().html('Uploading...please wait')
		    },
	        success: function(response){
				$("#imageloc").val(response.message)
				$("#storeblogimg").empty().html('Save Image')
				.removeClass("btn-info btn-danger").addClass("btn-success")
	            alert("Image Saved Successfully")
	        },
			error:function(response){
				$("#storeblogimg").empty().html('Save Image')
				.removeClass("btn-info btn-success").addClass("btn-danger")
			}
	    });
	});
	
	$("#storeimg").click(function(event){
		$(this).empty().html('Uploading...please wait')
		var pid = document.getElementById("pid")
		var guid = document.getElementById("guid")
		var data = new FormData();
	    $.each($('#imagehomepage')[0].files, function(i, file) {
	        data.append('file-'+i, file);
	    });
		data.append('pid',$(pid).val())
		data.append('guid', $(guid).val())
		$.ajax({
	        url: '/sendImageUploadimage',
			type: 'POST',
	        data: data,
	        cache: false,
	        contentType: false,
	        processData: false,
			before: function() {
			    $("#storeimg").empty().html('Uploading...please wait')
		    },
	        success: function(response){
	            alert("Image Saved Successfully")
				$("#storeimg").empty().html('Save Image')
				.removeClass("btn-info btn-danger").addClass("btn-success")
	        },
			error:function(response){
				console.log(response)
				$("#storeimg").empty().html('Save Image')
				.removeClass("btn-info btn-success").addClass("btn-danger")
			}
	    });
	});
	
	$("#storeshareimg").click(function(event){
		$(this).empty().html('Uploading...please wait')
		var pid = document.getElementById("pid")
		var guid = document.getElementById("guid")
		var data = new FormData();
	    $.each($('#imageshare')[0].files, function(i, file) {
	        data.append('fileNameShare', file);
	    });
		data.append('pid',$(pid).val())
		data.append('guid', $(guid).val())
		$.ajax({
	        url: '/sendBlogImageForShare',
			type: 'POST',
	        data: data,
	        cache: false,
	        contentType: false,
	        processData: false,
			before: function() {
			    $("#storeshareimg").empty().html('Uploading...please wait')
		    },
	        success: function(response){
				$("#imageloc").val(response.message)
				$("#storeshareimg").empty().html('Save Image')
				.removeClass("btn-info btn-danger").addClass("btn-success")
	            alert("Image Saved Successfully")
	        },
			error:function(response){
				$("#storeshareimg").empty().html('Save Image')
				.removeClass("btn-info btn-success").addClass("btn-danger")
			}
	    });
	});
	
	$('.btnShare').click(function(){
		postToFeed();
		return false;
	});
	
	$(".appr").click(function(elem){
		var source = $(this).attr("id");
		var f = source.split("-");
		var domID = f[1];
		var c = domID.split("_")
		var reviewID = c[1];
		var productID = c[2];
		var user = c[3];
		$.ajax(
           {
               url : '/api/approveRating',
               type: "POST",
               data: {reviewid: reviewID, pid: productID, user: user},
               beforeSend: function ()
               {
                 
               },
               success:function(response)
               {
                	
               	if(response.data == "success") {
						var currentV = parseInt(document.getElementById(domID).innerHTML)
						currentV += 1;
						document.getElementById(domID).innerHTML = currentV
					} else {
						alert(response.responseJSON.data)
					}
               },
               error: function(response)
               {
                	
               	if(response.status == 401) {
						alert("Sorry, you must be logged in to upvote")
					} else {
						alert(response.status + " " + response.statusText)
					}
               }
             });
	});
	
	$(".disappr").click(function() {
		var source = $(this).attr("id");
		var f = source.split("-");
		var domID = f[1];
		var c = domID.split("_")
		var reviewID = c[1];
		var productID = c[2];
		var user = c[3];
		$.ajax(
            {
                url : '/api/disapproveRating',
                type: "POST",
                data: {reviewid: reviewID, pid: productID, user: user},
                beforeSend: function ()
                {
                 
                },
                success:function(response)
                {
                	
                	if(response.data == "success") {
						var currentV = parseInt(document.getElementById(domID).innerHTML)
						currentV += 1;
						document.getElementById(domID).innerHTML = currentV
					} else {
						alert(response.data)
					}
                },
                error: function(response)
                {
                  	if(response.status == 401) {
						alert("Sorry, you must be logged in to downvote")
					} else {
						alert(response.status + " " + response.statusText)	
					}
                }
              });
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
