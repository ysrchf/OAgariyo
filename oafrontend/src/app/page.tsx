"use client"

import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import axios from "axios";
import {useForm} from "react-hook-form"


const Home =() =>{
    const queryClient = useQueryClient()

    const { register, handleSubmit, formState: { errors } } = useForm()

    const useDelete = () => {
        return useMutation({
            mutationFn: async (id: number) => {
                const {data} = await axios.delete("http://localhost:8080/v1/recipe/" + id)
                return data
            },
            onSuccess: () => queryClient.invalidateQueries('getRecipe')
        })
    }

    const usePost = () => {
        return useMutation({
            mutationFn: async (body: any) => {
                const {data} = await axios.post("http://localhost:8080/v1/recipe", body)
                return data
            },
            onSuccess: () => queryClient.invalidateQueries('getRecipe')
        })
    }

    const useGet = () => {
        return useQuery({
            queryKey: ["getRecipe"],
            queryFn: async () => {
                const {data} = await axios.get("http://localhost:8080/v1/recipe")
                return data
            },
        })
    }
    const { data, isLoading } = useGet()
    const {mutate: create, isPending: isLoadpost} = usePost()
    const {mutate: del, isPending: isLoaddel} = useDelete()

    if (isLoading || isLoadpost || isLoaddel) {
        return <div>loading</div>
    }

    const handleClick = () =>{
        console.log(data)
    }

    interface ResultItem{
        Id: number,
        Name: String,
        Ingredients: String,
        Rating: number
    }



    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
                {data.map((recipe: ResultItem ) => (
                    <div key={recipe.Id}>
                        <h3>{recipe.Name}</h3>
                        <div>{recipe.Ingredients}</div>
                        <div>{recipe.Rating}</div>
                        <button onClick={() => del(recipe.Id)}>Delete</button>
                    </div>
                ))}
            </div>

            <div>
                <form onSubmit={handleSubmit((rdata) => create(rdata))}>
                    <label htmlFor="name">Name</label>
                    <input type="text" placeholder="Soupe de Choux" {...register('Name')} />

                    <label htmlFor="ingredients">Ingredients</label>
                    <input type="text" placeholder="Choux, Eau, etc.." {...register('Ingredients')} />


                    <button type="submit">Créer Recette</button>
                </form>
            </div>

        </main>
    );

}

export default Home