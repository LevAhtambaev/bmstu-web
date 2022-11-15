import React from 'react'
import Slider from "@mui/material/Slider";
import Box from "@mui/material/Box";

export function Filter() {
    return (
        <div className="bg-blend-color-burn">
            <Box sx={{ width: 300 }}>
                <Slider
                    aria-label="Price filter"
                    defaultValue={0}
                    valueLabelDisplay="auto"
                    step={50}
                    marks
                    min={0}
                    max={1200}
                />
            </Box>
        </div>
    )
}