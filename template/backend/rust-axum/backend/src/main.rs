use axum::{
    routing::get,
    Router,
};

#[tokio::main]
async fn main() {
    let app = Router::new().route("/", get(|| async { "Hello, World!" }));

    let listener = tokio::net::TcpListener::bind("{{ .Address }}:{{ .Port }}").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
