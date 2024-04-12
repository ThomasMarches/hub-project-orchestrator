use actix_web::{web, App, HttpResponse, HttpServer};

async fn index() -> HttpResponse {
    HttpResponse::Ok().body("Hello world!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().route("/", web::get().to(index)))
        .bind(("{{ .Address }}", {{ .Port }}))?
        .run()
        .await
}
