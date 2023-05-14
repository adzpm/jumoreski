(function() {
    // Add event listener
    document.addEventListener("mousemove", parallax);
    const elem = document.querySelector("#parallax");
    // Magic happens here
    function parallax(e) {
        let _w = window.innerWidth/2;
        let _h = window.innerHeight/2;
        let _mouseX = e.clientX;
        let _mouseY = e.clientY;
        let _depth1 = `${50 - (_mouseX - _w) * 0.01}% ${50 - (_mouseY - _h) * 0.01}%`;
        let _depth2 = `${50 - (_mouseX - _w) * 0.02}% ${50 - (_mouseY - _h) * 0.02}%`;
        let _depth3 = `${50 - (_mouseX - _w) * 0.06}% ${50 - (_mouseY - _h) * 0.06}%`;
        let x = `${_depth3}, ${_depth2}, ${_depth1}`;
        console.log(x);
        elem.style.backgroundPosition = x;
    }

})();

axios.defaults.withCredentials = true

let Application = Vue.createApp({
    data: function () {
        return {
            anek: '',
        }
    },

    methods: {
        formatSpaces: function (text) {
            return text.replaceAll(/\r?\n/g, `<br>`)
        },

        getRandomAnek: function() {
            axios.get('http://127.0.0.1:8080/random').then(
                response => {
                    this.anek = response.data.text
                }
            )
        }
    },

    mounted: function () {
    },

    computed: {},
})

let wm = Application.mount('#app')