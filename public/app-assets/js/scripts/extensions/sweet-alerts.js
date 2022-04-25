/*=========================================================================================
	File Name: sweet-alerts.js
	Description: A beautiful replacement for javascript alerts
	----------------------------------------------------------------------------------------
	Item Name: Stack - Responsive Admin Theme
	Version: 1.0
	Author: GeeksLabs
	Author URL: http://www.themeforest.net/user/geekslabs
==========================================================================================*/
$(document).ready(function(){

	$('#basic-alert').on('click',function(){
		swal("Here's a message!");
	});

	$('#with-title').on('click',function(){
		swal("Here's a message!", "It's pretty, isn't it?");
	});

	$('#html-alert').on('click',function(){
		swal({
			title: 'HTML <small>Title</small>!',
			text: 'A custom <span style="color:#F6BB42">html<span> message.',
			html: true
		});
	});

	$('#type-success').on('click',function(){
		swal("Good job!", "You clicked the button!", "success");
	});

	$('#type-info').on('click',function(){
		swal("Info!", "You clicked the button!", "info");
	});

	$('#type-warning').on('click',function(){
		swal("Warning!", "You clicked the button!", "warning");
	});

	$('#type-error').on('click',function(){
		swal("Error!", "You clicked the button!", "error");
	});

	$('#custom-icon').on('click',function(){
		swal({   title: "Sweet!",   text: "Here's a custom image.",   imageUrl: "../../../app-assets/images/icons/thumbs-up.jpg" });
	});

	$('#auto-close').on('click',function(){
		swal({   title: "Auto close alert!",   text: "I will close in 2 seconds.",   timer: 2000,   showConfirmButton: false });
	});

	$('#outside-click').on('click',function(){
		swal({
			title: 'Click outside to close!',
			text: 'This is a cool message!',
			allowOutsideClick: true
		});
	});

	$('#cancel-button').on('click',function(){
		swal({
		    title: "Are you sure?",
		    text: "You will not be able to recover this imaginary file!",
		    type: "warning",
		    showCancelButton: true,
		    confirmButtonColor: "#F6BB42",
		    confirmButtonText: "Yes, delete it!",
		    cancelButtonText: "No, cancel plx!",
		    closeOnConfirm: false,
		    closeOnCancel: false
		}, function(isConfirm) {
		    if (isConfirm) {
		        swal("Deleted!", "Your imaginary file has been deleted.", "success");
		    } else {
		        swal("Cancelled", "Your imaginary file is safe :)", "error");
		    }
		});

	});

	$('#prompt-function').on('click',function(){
		swal({
		    title: "An input!",
		    text: "Write something interesting:",
		    type: "input",
		    showCancelButton: true,
		    closeOnConfirm: false,
		    animation: "slide-from-top",
		    inputPlaceholder: "Write something"
		}, function(inputValue) {
		    if (inputValue === false) return false;
		    if (inputValue === "") {
		        swal.showInputError("You need to write something!");
		        return false
		    }
		    swal("Nice!", "You wrote: " + inputValue, "success");
		});

	});

	$('#ajax-request').on('click',function(){
		swal({
		    title: "Ajax request example",
		    text: "Submit to run ajax request",
		    type: "info",
		    showCancelButton: true,
		    closeOnConfirm: false,
		    showLoaderOnConfirm: true,
		}, function() {
		    setTimeout(function() {
		        swal("Ajax request finished!");
		    }, 2000);
		});
	});

	$('#confirm-text').on('click',function(){
		swal({
		    title: "Confirm Button Text",
		    text: "See the confirm button text! Have you noticed the Change?",
		    type: "warning",
		    showCancelButton: true,
		    confirmButtonText: "Text Changed!!!",
		    cancelButtonText: "No, cancel plx!",
		    closeOnConfirm: false,
		    closeOnCancel: false
		}, function(isConfirm) {
		    if (isConfirm) {
		        swal("Changed!", "Confirm button text was changed!!", "success");
		    } else {
		        swal("Cancelled", "It's safe.", "error");
		    }
		});
	});

	$('#confirm-color').on('click',function(){
		swal({
		    title: "Are you sure?",
		    text: "You will not be able to recover this imaginary file!",
		    type: "warning",
		    showCancelButton: true,
		    confirmButtonColor: "#DA4453",
		    confirmButtonText: "Yes, delete it!",
		    cancelButtonText: "No, cancel plx!",
		    closeOnConfirm: false,
		    closeOnCancel: false
		}, function(isConfirm) {
		    if (isConfirm) {
		        swal("Deleted!", "Your imaginary file has been deleted.", "success");
		    } else {
		        swal("Cancelled", "Your imaginary file is safe :)", "error");
		    }
		});
	});

	$('#pop-animation').on('click',function(){
		swal({
			title: 'Default Animation',
			text: "POP Animation!",
			animation: "pop",
		});
	});

	$('#slide-top-animation').on('click',function(){
		swal({
			title: 'Slide Animation',
			text: "Slide From Top Animation",
			animation: "slide-from-top",
		});

	});

	$('#slide-bottom-animation').on('click',function(){
		swal({
			title: 'Slide Animation',
			text: "Slide From Bottom Animation",
			animation: "slide-from-bottom",
		});
	});

});