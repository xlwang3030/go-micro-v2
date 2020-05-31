<template>
  <nav class="navbar is-light" role="navigation" aria-label="main navigation">
    <div class="container">
      <div class="navbar-brand">
      <a class="navbar-item " href="http://www.jtthink.com">
        <img src="/logox.png"  />
      </a>
      <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>
      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item"  href="/">
            首页
          </a>
          <a class="navbar-item" href="/course">
            推荐
          </a>

          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">
              分类
            </a>
            <div class="navbar-dropdown">
              <a class="navbar-item">
                Go专场
              </a>
              <a class="navbar-item">
                Java专场
              </a>
              <a class="navbar-item">
                PHP专场
              </a>
              <hr class="navbar-divider">
              <a class="navbar-item">
                社区
              </a>
            </div>
          </div>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <div class="buttons">
              <a class="button is-primary">
                <strong>注册</strong>
              </a>
              <a class="button is-light">
                登录
              </a>
            </div>
          </div>
        </div>
      </div></div>

  </nav>
</template>
<script>
  export default {
    mounted() {
      document.addEventListener('DOMContentLoaded', () => {

        // Get all "navbar-burger" elements
        const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

        // Check if there are any navbar burgers
        if ($navbarBurgers.length > 0) {

          // Add a click event on each of them
          $navbarBurgers.forEach( el => {
            el.addEventListener('click', () => {

              // Get the target from the "data-target" attribute
              const target = el.dataset.target;
              const $target = document.getElementById(target);

              // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
              el.classList.toggle('is-active');
              $target.classList.toggle('is-active');

            });
          });
        }

      });
    }
  }
</script>
