$("#login").on('submit', login)

function login(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $("#email").val(),
            password: $("#password").val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Invalid email or password!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        })
    });
}