
$(document).ready(function() {
    /********* On scroll heder Sticky *********/
    $(window).scroll(function() {
        var scroll = $(window).scrollTop();
        if (scroll >= 50) {
            $("header").addClass("head-sticky");
        } else {
            $("header").removeClass("head-sticky");
        }
    });   
    /********* Mobile Menu ********/  
   
    /********* Cart Popup ********/
    
    /********* Mobile Filter Popup ********/
    $('.filter-title').on('click',function(e){
        e.preventDefault();
        setTimeout(function(){
            $('body').addClass('no-scroll filter-open');
            $('.overlay').addClass('active');
        }, 50);
    }); 
    $('body').on('click','.overlay.active, .close-filter', function(e){
        e.preventDefault(); 
        $('.overlay').removeClass('active');
        $('body').removeClass('no-scroll filter-open');
    });
    /* hiện thanh tìm kiếm SEARCH */ 
    $(".search-header a").click(function() { 
        $(".search-popup").toggleClass("active"); 
        $("body").toggleClass("no-scroll");
    });
    $(".close-search").click(function() { 
        $(".search-popup").removeClass("active"); 
        $("body").removeClass("no-scroll");
    });
    /******* Cookie Js *******/
    $('.cookie-close').click(function () {
        $('.cookie').slideUp();
    });
    /******* Subscribe popup Js *******/
    $('.close-sub-btn').click(function () {
        $('.subscribe-popup').slideUp(); 
        $(".subscribe-overlay").removeClass("open");
    });      
    /********* qty spinner ********/
    var quantity = 0;
    $('.quantity-increment').click(function(){;
        var t = $(this).siblings('.quantity');
        var quantity = parseInt($(t).val());
        $(t).val(quantity + 1); 
    }); 
    $('.quantity-decrement').click(function(){
        var t = $(this).siblings('.quantity');
        var quantity = parseInt($(t).val());
        if(quantity > 1){
            $(t).val(quantity - 1);
        }
    });   
    /******  Nice Select  ******/ 
    $('select').niceSelect(); 
    /*********  Multi-level accordion nav  ********/ 
    $('.acnav-label').click(function () {
        var label = $(this);
        var parent = label.parent('.has-children');
        var list = label.siblings('.acnav-list');
        if (parent.hasClass('is-open')) {
            list.slideUp('fast');
            parent.removeClass('is-open');
        }
        else {
            list.slideDown('fast');
            parent.addClass('is-open');
        }
    }); 
    /****  TAB Js ****/
    $('ul.tabs li').click(function(){ 
        var tab_id = $(this).attr('data-tab'); 
        $('ul.tabs li').removeClass('active');
        $('.tab-content').removeClass('active');  
        $(this).addClass('active');
        $("#"+tab_id).addClass('active');
        $('.bestsell-slider').slick('refresh');
	}) 
    if($('.client-logo-slider').length > 0 ){
        $('.client-logo-slider').slick({
            autoplay: true, 
            slidesToShow: 5,
            speed: 1000,
            slidesToScroll: 1,  
            prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
            nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
            dots: false,
            buttons: false,
            responsive: [ 
                {
                    breakpoint: 1200,
                    settings: {
                    slidesToShow: 4,
                    slidesToScroll: 1   
                    }
                },  
                {
                breakpoint: 992,
                    settings: {
                        slidesToShow: 3,
                        slidesToScroll: 1 
                    }
                },
                {
                breakpoint: 576,
                    settings: {
                        slidesToShow: 1,
                        slidesToScroll: 1 
                    }
                }
            ]
        });
    }  
    $('.hiro-image-slider').slick({ 
          arrows: false,
          dots: false,
          speed: 500,
          slidesToShow: 1,
          infinite: false,
          cssEase: 'ease-in-out',
          swipe: false,
          fade: true,
          slidesToScroll: 1,  
          asNavFor: '.hiro-side-slider, .hiro-thumb-slider'
      });
      $('.hiro-side-slider').slick({ 
          arrows: false,
          fade: true, 
          infinite: false, 
          speed: 500,
          swipe: false,
          cssEase: 'ease-in-out',
          slidesToShow: 1,
          slidesToScroll: 1,  
          adaptiveHeight: true,
          asNavFor: '.hiro-thumb-slider, .hiro-image-slider'
      });
      $('.hiro-thumb-slider').slick({ 
           arrows: false, 
           infinite: false,
           slidesToShow: 4,   
           slidesToScroll: 1,    
           asNavFor: '.hiro-side-slider, .hiro-image-slider',
           focusOnSelect: true,
      });
    $('.playbutton-sec').click(function (m) { 
        $('body,html').addClass('no-scroll popupopen');
        $('.overlay-popup').addClass('popup-show');
    });  
    $('.close-popup').click(function (m) { 
        $('body,html').removeClass('no-scroll popupopen');  
        $('.overlay-popup').removeClass('popup-show');
    });    
    
    
    /** PDP slider **/
    $('.two-coll-content .thmb-pro-main').slick({
        dots: false,
        infinite: false,
        speed: 500,
        loop: false,
        slidesToShow: 1,
        cssEase: 'ease-in-out',
        swipe: false,
        fade: true,
        arrows: false,
        asNavFor: '.two-coll-content .thumb-pro-list',
    });
    $('.two-coll-content .thumb-pro-list').slick({
        arrows: false,
        asNavFor: '.two-coll-content .thmb-pro-main', 
        speed: 500,
        dots: false,
        slidesToScroll: 1, 
        loop: false,
        infinite: false, 
        focusOnSelect: true,
        slidesToShow: 3
    });
    
    $('.all-shoes-slider').slick({
        autoplay: false, 
        slidesToShow: 3,
        speed: 1000,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [  
            {
            breakpoint:992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 481,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    $('.bestsell-slider').slick({
        autoplay: false, 
        slidesToShow: 2,
        speed: 1000,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 768,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    $('.all-shes-second-slider').slick({
        autoplay: false, 
        slidesToShow: 4,
        speed: 1000,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 1200,
                settings: {
                    slidesToShow: 3,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 591,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    $('.testimonial-slider').slick({
        autoplay: false, 
        slidesToShow: 4,
        speed: 1000,
        infinite:false,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 1200,
                settings: {
                    slidesToShow: 3,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 591,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    $('.product-card-slider').slick({
        autoplay: false, 
        slidesToShow: 4,
        speed: 1000,
        infinite:false,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 1200,
                settings: {
                    slidesToShow: 3,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 420,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    });   

    // chỉnh phần Realted Products
    $('.related-product-slider').slick({
        autoplay: false, 
        slidesToShow: 6,
        speed: 1000,
        infinite:false,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 1200,
                settings: {
                    slidesToShow: 3,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 420,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    $('.our-blogs-slider').slick({
        autoplay: false, 
        slidesToShow: 4,
        speed: 1000,
        infinite:false,
        slidesToScroll: 1,  
        prevArrow: '<button class="slick-prev slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        nextArrow: '<button class="slick-next slick-arrow"><span class="slickbtn"><svg><use xlink:href="#slickarrow"></use></svg></span></button>',
        dots: false,
        buttons: false,
        responsive: [   
            {
            breakpoint: 1200,
                settings: {
                    slidesToShow: 3,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 992,
                settings: {
                    slidesToShow: 2,
                    slidesToScroll: 1 
                }
            },
            {
            breakpoint: 420,
                settings: {
                    slidesToShow: 1,
                    slidesToScroll: 1 
                }
            }
        ]
    }); 
    /** PDP slider **/
    $('.pro-main-slider').slick({
        dots: false,
        infinite: true,
        speed:500,
        loop: true,
        slidesToShow: 1,
        arrows: false,
        asNavFor: '.pdp-thumb-slider',
    });
    $('.pdp-thumb-slider').slick({
        prevArrow: '<button class="slide-arrow slick-prev"><svg viewBox="0 0 10 5"><path d="M2.37755e-08 2.57132C-3.38931e-06 2.7911 0.178166 2.96928 0.397953 2.96928L8.17233 2.9694L7.23718 3.87785C7.07954 4.031 7.07589 4.28295 7.22903 4.44059C7.38218 4.59824 7.63413 4.60189 7.79177 4.44874L9.43039 2.85691C9.50753 2.78197 9.55105 2.679 9.55105 2.57146C9.55105 2.46392 9.50753 2.36095 9.43039 2.28602L7.79177 0.69418C7.63413 0.541034 7.38218 0.544682 7.22903 0.702329C7.07589 0.859976 7.07954 1.11192 7.23718 1.26507L8.1723 2.17349L0.397965 2.17336C0.178179 2.17336 3.46059e-06 2.35153 2.37755e-08 2.57132Z"></path></svg></button>',
        nextArrow: '<button class="slide-arrow slick-next"><svg viewBox="0 0 10 5"><path d="M2.37755e-08 2.57132C-3.38931e-06 2.7911 0.178166 2.96928 0.397953 2.96928L8.17233 2.9694L7.23718 3.87785C7.07954 4.031 7.07589 4.28295 7.22903 4.44059C7.38218 4.59824 7.63413 4.60189 7.79177 4.44874L9.43039 2.85691C9.50753 2.78197 9.55105 2.679 9.55105 2.57146C9.55105 2.46392 9.50753 2.36095 9.43039 2.28602L7.79177 0.69418C7.63413 0.541034 7.38218 0.544682 7.22903 0.702329C7.07589 0.859976 7.07954 1.11192 7.23718 1.26507L8.1723 2.17349L0.397965 2.17336C0.178179 2.17336 3.46059e-06 2.35153 2.37755e-08 2.57132Z"></path></svg></button>',
        dots: false,
        asNavFor: '.pro-main-slider', 
        speed: 500, 
        slidesToScroll: 1,
        touchMove: true,
        focusOnSelect: true,
        loop: true,
        infinite: true,
        focusOnSelect: true,
        slidesToShow: 4,
        responsive: [{
                breakpoint: 1261,
                settings: {
                    slidesToShow: 3 
                }
            } 
        ]   
    }); 
    //video-play 
    $('.play-vid').on('click',function(){
        if($(this).attr('data-click') == 1) {
        $(this).attr('data-click', 0)
        $('#img-vid')[0].pause();
        $(".play-vid").css("opacity", "1");
        } else {
        $(this).attr('data-click', 1)
        $('#img-vid')[0].play();
        $(".play-vid").css("opacity", "1");
        $(".play-vid").css("opacity", "0");
        }
    });
});  
if ($(".my-acc-column").length > 0) {
    jQuery(function ($) {
        var topMenuHeight = $("#daccount-nav").outerHeight();
        $("#account-nav").menuScroll(topMenuHeight);
        $(".account-list li:first-child").addClass("active");
    });
    // COPY THE FOLLOWING FUNCTION INTO ANY SCRIPTS
    jQuery.fn.extend({
        menuScroll: function (offset) {
            // Declare all global variables
            var topMenu = this;
            var topOffset = offset ? offset : 0;
            var menuItems = $(topMenu).find("a");
            var lastId;
            // Save all menu items into scrollItems array
            var scrollItems = $(menuItems).map(function () {
                var item = $($(this).attr("href"));
                if (item.length) {
                    return item;
                }
            });
            // When the menu item is clicked, get the #id from the href value, then scroll to the #id element
            $(topMenu).on("click", "a", function (e) {
                var href = $(this).attr("href");
                var offsetTop = href === "#" ? 0 : $(href).offset().top - topOffset;
                function checkWidth() {
                    var windowSize = $(window).width();
                    if (windowSize <= 767) {
                        $('html, body').stop().animate({
                            scrollTop: offsetTop - 200
                        }, 300);
                    }
                    else {
                        $('html, body').stop().animate({
                            scrollTop: offsetTop - 100
                        }, 300);
                    }
                }
                // Execute on load
                checkWidth();
                // Bind event listener
                $(window).resize(checkWidth);
                e.preventDefault();
            });
            // When page is scrolled
            $(window).scroll(function () {
                function checkWidth() {
                    var windowSize = $(window).width();
                    if (windowSize <= 767) {
                        var nm = $("html").scrollTop();
                        var nw = $("body").scrollTop();
                        var fromTop = (nm > nw ? nm : nw) + topOffset;
                        // When the page pass one #id section, return all passed sections to scrollItems and save them into new array current
                        var current = $(scrollItems).map(function () {
                            if ($(this).offset().top - 250 <= fromTop)
                                return this;
                        });
                        // Get the most recent passed section from current array
                        current = current[current.length - 1];
                        var id = current && current.length ? current[0].id : "";
                        if (lastId !== id) {
                            lastId = id;
                            // Set/remove active class
                            $(menuItems)
                                .parent().removeClass("active")
                                .end().filter("[href='#" + id + "']").parent().addClass("active");
                        }
                    }
                    else {
                        var nm = $("html").scrollTop();
                        var nw = $("body").scrollTop();
                        var fromTop = (nm > nw ? nm : nw) + topOffset;
                        // When the page pass one #id section, return all passed sections to scrollItems and save them into new array current
                        var current = $(scrollItems).map(function () {
                            if ($(this).offset().top <= fromTop)
                                return this;
                        });
                        // Get the most recent passed section from current array
                        current = current[current.length - 1];
                        var id = current && current.length ? current[0].id : "";
                        if (lastId !== id) {
                            lastId = id;
                            // Set/remove active class
                            $(menuItems)
                                .parent().removeClass("active")
                                .end().filter("[href='#" + id + "']").parent().addClass("active");
                        }
                    }
                }
                // Execute on load
                checkWidth();
                // Bind event listener
                $(window).resize(checkWidth);
            });
        }
    });
}
$(window).on('load resize orientationchange', function() { 
    /********* Wrapper top space ********/
    var header_hright = $('header').outerHeight();
    $('header').next('.wrapper').css('margin-top', header_hright + 'px');  
}); 