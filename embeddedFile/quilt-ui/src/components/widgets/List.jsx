import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import parse from 'html-react-parser';

import { get, map } from 'lodash';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
    },
    header: {
        marginLeft: "0.5rem",
        ...typography.heading1
    },
    tileTextContainer: {
        margin: "0.75rem",
    },
    tileImage: {
        maxWidth: "100%"
    },
    tileTextTitle: {
        ...typography.heading2,
        marginBottom: "0.5rem"
    },
    tileTextSubtitle: {
        ...typography.body1,
        lineHeight: "1",
        color: colors.text02Color,
    },
    tileHtmlTextContainer: {
        marginRight: "0.75rem",
    },
}));

const List = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [components, setComponents] = useState([]);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        if (shouldRenderEditorComponent && widgetData.length != 0) {
            // updateLocalisationValueFunc is a wrapper function for updateLocalisationValue to avoid having to pass in params(that are local to this scope) to renderEditorFieldComponents function.
            const updateLocalisationValueFunc = (localeKey, localeValue) => {
                updateLocalisationValue({localeKey, localeValue, pageId, setErrorMessage, setSaving, setEditing, fetchPageData})
            }
            // Generate EditorFieldComponents based on widgetData
            const editorFieldComponents = renderEditorFieldComponents({widgetData, saving, setWidgetData, editing, setEditing, initState, updateLocalisationValueFunc})
            // Invoke updateEditorComponent to render the EditorComponent
            updateEditorComponent(
                <EditorContainer 
                    errorMessage={errorMessage} 
                    editorFieldComponents={editorFieldComponents}
                    resetFunction={resetEditorComponent}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage])

    // Render Preview component
    useEffect(() => {
        setComponents(
            map(widgetData, (value, type) => {
                if(type.includes("tile-image")) {
                    const imageUrlParsedValue = getValueFromWidgetData(widgetData, type, "data.image_url")
                    return <img 
                        key={type} 
                        src={imageUrlParsedValue}
                        alt="tooltip" 
                        className={clsx(classes.tileImage)} 
                    />
                } else if (type.includes("tile-text")) {
                    const titleParsedValue = getValueFromWidgetData(widgetData, type, "data.title")
                    const subtitleParsedValue = getValueFromWidgetData(widgetData, type, "data.subtitle")
                    return <Box 
                        key={type}
                        className={clsx(classes.tileTextContainer)}
                    >
                        <Typography 
                            variant="body1"
                            className={clsx(classes.tileTextTitle)}
                        >
                            {titleParsedValue}
                        </Typography>
                        <Typography 
                            variant="body1"
                            className={clsx(classes.tileTextSubtitle)}
                        >
                            {subtitleParsedValue}
                        </Typography>
                    </Box>
                } else if (type.includes("tile-html-text")) {
                    const titleParsedValue = getValueFromWidgetData(widgetData, type, "data.title")
                    return <Box 
                        key={type}
                        className={clsx(classes.tileHtmlTextContainer)}
                    >
                        {parse(titleParsedValue)}
                    </Box>
                }
            })
        )
    }, [widgetData])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget || get(widget, "components.length") === 0) return
        setKeyValueToState(widget, "header", "header.title", setWidgetData)
        for(let i=0; i<widget.components.length; i++) {
            const component = widget.components[i]
            switch(component.type) {
                case "tile-image" : {
                    setKeyValueToState(component, `tile-image-${i}`, "data.image_url", setWidgetData)
                    continue
                }
                case "tile-text" : {
                    setKeyValueToState(component, `tile-text-${i}`, "data.title", setWidgetData)
                    setKeyValueToState(component, `tile-text-${i}`, "data.subtitle", setWidgetData)
                    continue
                }
                case "tile-html-text" : {
                    setKeyValueToState(component, `tile-html-text-${i}`, "data.title", setWidgetData)
                    continue
                }
            }
        }
    }

    useEffect(() => {
        initState()
    }, [widget])

    if (get(widget, "components.length") ===0) {
        return <div>Empty {widget.type}</div>
    }

    const headerTitleParsedValue = getValueFromWidgetData(widgetData, "header", "header.title")
    return (widgetData && <Box className={clsx(classes.widgetContainer)}>
        <Box>
            <Typography 
                className={clsx(classes.header)} 
                variant="body1">
                    {headerTitleParsedValue}
            </Typography>
        </Box>
        <Grid container direction="row">
            {map(components, component => component)}
        </Grid>
    </Box>)
}

export default List;