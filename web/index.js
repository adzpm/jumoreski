// (function() {
//     // Add event listener
//     document.addEventListener("mousemove", parallax);
//     const elem = document.querySelector("#parallax");
//     // Magic happens here
//     function parallax(e) {
//         let _w = window.innerWidth/2;
//         let _h = window.innerHeight/2;
//         let _mouseX = e.clientX;
//         let _mouseY = e.clientY;
//         let _depth1 = `${50 - (_mouseX - _w) * 0.01}% ${50 - (_mouseY - _h) * 0.01}%`;
//         let _depth2 = `${50 - (_mouseX - _w) * 0.02}% ${50 - (_mouseY - _h) * 0.02}%`;
//         let _depth3 = `${50 - (_mouseX - _w) * 0.06}% ${50 - (_mouseY - _h) * 0.06}%`;
//         let x = `${_depth3}, ${_depth2}, ${_depth1}`;
//         console.log(x);
//         elem.style.backgroundPosition = x;
//     }
//
// })();

axios.defaults.withCredentials = true

let Application = Vue.createApp({
    data: function () {
        return {
            loading: false,
            obj: {
                text: "Hello, world!",
                images: [
                    "https://media.tenor.com/CeklkfKEOnAAAAAM/gachi-gachimuchi-billy-harington-chmok.gif",
                    "https://backupuusites.hb.bizmrg.com/resize_cache/6096307/d5cb5488396720686a69d0c49ef80752/iblock/d45/d45a96c13aa3585ed2e809753cb0b63d/c8de51edbde3e6b51dad851eab7e45f1.jpg",
                ],
            },
            loaders : [
                "https://steamuserimages-a.akamaihd.net/ugc/941715990349187757/FFACE0D31D7EDBBE641F0C31AAFF835BA7090F03/?imw=5000&imh=5000&ima=fit&impolicy=Letterbox&imcolor=%23000000&letterbox=false",
                "https://media.tenor.com/8R2DqGM5GMwAAAAd/gachi-billy-herrington.gif",
                "https://media.tenor.com/3CrIwgbopCUAAAAM/billy-gachi.gif",
                "https://media.tenor.com/6V1ooQVn3CQAAAAC/gachi-fist.gif",
                "https://media.tenor.com/okJxD0a2jaYAAAAM/gachi-men.gif",
                "https://media.tenor.com/ud0bG4WnWRgAAAAC/gachi-gachimuchi.gif",
                "https://media.tenor.com/J01NX0ikmNwAAAAC/ricardo-flick-dance.gif",
            ]
        }
    },

    methods: {
        formatSpaces: function (text) {
            return text.replaceAll(/\r?\n/g, `<br>`)
        },

        getRandomAnek: function() {
            this.loading = true
            axios.get('http://127.0.0.1:8080/random').then(
                response => {
                    this.obj.text = response.data.text
                    this.obj.images = response.data.images
                }
            ).finally(_ => {
                setTimeout(_ => {
                    this.loading = false
                }, 1000)
            })
        },

        getRandomLoadingImage: function() {
            return this.loaders[Math.floor(Math.random() * this.loaders.length)]
        }
    },

    mounted: function () {
    },

    computed: {},
})

let wm = Application.mount('#app')