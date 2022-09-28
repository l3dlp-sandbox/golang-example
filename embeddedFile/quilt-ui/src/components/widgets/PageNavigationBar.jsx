import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        marginTop: "1rem",
        marginBottom: "1rem",
    },
    back: {
        marginLeft: "0.5rem",
        ...typography.heading3
    },
    headerTitle: {
        ...typography.body1,
        textAlign: "center"
    }
}));

const PageNavigationBar = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
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

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return
        setKeyValueToState(widget, "header", "header.title", setWidgetData)
    }

    useEffect(() => {
        initState()
    }, [widget])

    const headerTitle = getValueFromWidgetData(widgetData, "header", "header.title")
    return (<Box className={clsx(classes.widgetContainer)}>
        <Grid container direction="row">
            <Grid item xs={2}>
                <Typography 
                    className={clsx(classes.back)} 
                    variant="body1">
                        Back
                </Typography>
            </Grid>
            <Grid item xs={10}>
                <Typography 
                    className={clsx(classes.headerTitle)} 
                    variant="body1">
                        {headerTitle}
                </Typography>
            </Grid>
        </Grid>
    </Box>)
}

export default PageNavigationBar;