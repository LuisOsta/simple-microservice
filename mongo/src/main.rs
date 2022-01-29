
fn user_generator(_history: &planter::GeneratedData) -> bson::Document {
    return bson::doc! {
        "username": "janedoe",
        "password": "placeholder"
    };
}

fn main() {
    let documents_per_collection: i32 = 1;
    let collections: Vec<(String, planter::EntityGenerator)> = vec![
        (String::from("auth"), user_generator),
    
    ];

    planter::seed_data(
        collections,
        planter::Configurations::new(documents_per_collection, planter::SeedMode::Disk),
    );

}
