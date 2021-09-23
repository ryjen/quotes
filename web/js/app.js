
const app = (function() {

  // private 
  const header = (function() {
    let nav = null

    const toggle = () => {
      nav.slideToggle({
          duration: ANIM_DELAY,
          start: function() {
            jQuery(this).css('display', 'flex');
          }
      });
    }

    const init = () => {
      nav = $("#header nav")
      menu = $("#menu-toggle")

      menu.click(menu_toggle)

      if (menu.is(":visible")) {
        nav.hide()
      }

      window.onresize = () => {
        if (menu.is(":visible")) {
          nav.hide()
        } else {
          nav.show()
        }
      }
    }
    return { init }
  }())

  const init = () => {
    header.init()
    console.log("app initialized")
  }

  // return public api
  return { 
    init
  }

}())

app.init()