import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { get } from 'lodash';

import { getValueFromWidgetData, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors, commonStyles } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    manageCard: {
        background: colors.general01Color,
        borderRadius: "10px",
        margin: "0.35rem",
        ...commonStyles.shadow
    },
    manageTitle: {
        ...typography.heading3,
        marginTop: "0.3rem",
        marginBottom: "0.2rem",
        paddingLeft: "0.5rem",
    },
    manageDescription: {
        ...typography.body2,
        lineHeight: "1",
        paddingLeft: "0.5rem",
    },
    manageImage: {
        minHeight: "100%",
        maxWidth: "2.5rem",
    }
}));

const GridMaxTwoExpandableComponent = ({widget, updateEditorComponent, resetSelectedComponentCard, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        if (shouldRenderEditorComponent && widgetData) {
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
                    resetFunction={resetSelectedComponentCard}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage])

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    const initState = () => {
        if (!widget) return
        setKeyValueToState(widget, "title", "data.title", setWidgetData)
        setKeyValueToState(widget, "description", "data.description", setWidgetData)
        setKeyValueToState(widget, "image", "data.image", setWidgetData)
    }

    useEffect(() => {
        initState()
    }, [widget])

    const parsedTitle = getValueFromWidgetData(widgetData, "title", "data.title")
    const parsedDescription = getValueFromWidgetData(widgetData, "description", "data.description")
    const parsedImage = getValueFromWidgetData(widgetData, "image", "data.image")
    const componentId = get(widget, "data.id")
    return (widgetData && <Grid item xs={6} key={componentId}>
        <Box className={clsx(classes.manageCard)}>
            <Grid container direction="row">
                <Grid item xs={3}>
                    <img src={parsedImage} alt="tooltip" className={clsx(classes.manageImage)} />
                </Grid>
                <Grid item xs={9}>
                    <Typography 
                        variant="body1"
                        className={clsx(classes.manageTitle)}
                    >
                        {parsedTitle}
                    </Typography>
                    <Typography 
                        variant="body1"
                        className={clsx(classes.manageDescription)}
                    >
                        {parsedDescription}
                    </Typography>
                </Grid>
            </Grid>
        </Box>
    </Grid>)
}

export default GridMaxTwoExpandableComponent;