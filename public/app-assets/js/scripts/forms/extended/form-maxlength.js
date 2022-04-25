/*=========================================================================================
	File Name: form-maxlength.js
	Description: Bootstrap-Maxlength uses a Twitter Bootstrap label to show a visual
		feedback to the user about the maximum length of the field where the user is
		inserting text. Uses the HTML5 attribute "maxlength" to work.
	----------------------------------------------------------------------------------------
	Item Name: Stack - Responsive Admin Theme
	Version: 1.0
	Author: GeeksLabs
	Author URL: http://www.themeforest.net/user/geekslabs
==========================================================================================*/
(function(window, document, $) {
	'use strict';
	// Default usage
	$('.basic-maxlength').maxlength({
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
	});

	// Change the threshold value
	$('.threshold-maxlength').maxlength({
		threshold: 15,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
	});

	// AlwaysShow
	$('.always-show-maxlength').maxlength({
		alwaysShow: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
	});

	// Change Badge Color using warningClass & limitReachedClass
	$('.badge-maxlength').maxlength({
		warningClass: "tag tag-info",
		limitReachedClass: "tag tag-warning"
	});

	// Change Badge Format
	$('.badge-text-maxlength').maxlength({
		alwaysShow: true,
		separator: ' of ',
		preText: 'You have ',
		postText: ' chars remaining.',
		validate: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
	});

	// Position
	$('.position-maxlength').maxlength({
		alwaysShow: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
		placement: 'top'
		// Options : top, bottom, left or right
		//  bottom-right, top-right, top-left, bottom-left and centered-right.
	});

	$('.position-corner-maxlength').maxlength({
		alwaysShow: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
		placement: 'top-left'
		//  bottom-right, top-right, top-left, bottom-left and centered-right.
	});

	$('.position-inside-maxlength').maxlength({
		alwaysShow: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
		placement: 'centered-right'
		// Option : centered-right.
	});

	$('.featured-maxlength').maxlength({
		alwaysShow: true,
		threshold: 10,
		warningClass: "tag tag-info",
		limitReachedClass: "tag tag-warning",
		placement: 'top',
		message: 'Used %charsTyped% of %charsTotal% chars.'
	});

	// Teatarea Maxlength
	$('.textarea-maxlength').maxlength({
		alwaysShow: true,
		warningClass: "tag tag-success",
		limitReachedClass: "tag tag-danger",
	});

})(window, document, jQuery);