import React, {useEffect, useState} from "react";
import { DataGrid, GridColDef } from "@mui/x-data-grid";

interface Album {
    id: number,
    title: string
}

function Albums() {
    const [data, setData] = useState<Album[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const renderAlbums = (data: Album[]) => {
        return data.map(d => <div>d.Title</div>);
    }

    useEffect(() => {
        (async () => {
            const albumsResponse = await fetch("http://localhost:8080/api/v1/albums");
            const { albums } = await albumsResponse.json();
            setLoading(false);
            setData(albums);
        })();
    }, []);

    if (data === null) {
        return <div></div>;
    }

    const columns: GridColDef[] = [
        { field: "artist", headerName: "Artist", width: 300 },
        { field: "title", headerName: "Title", width: 400 },
        { field: "price", headerName: "Price", width: 75 },
    ];

    return (
        <div className="Albums">
            {loading && <p>Loading...</p>}
            {!loading && <div style={{ height: 300, width: "100%" }}><DataGrid rows={data} columns={columns}/></div>}
        </div>
    );
}

export default Albums;