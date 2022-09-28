import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import { get, map } from 'lodash';

import AddonExpandableSection from './AddonExpandableSection';
import { parseQUIPrefix, getValueFromWidgetData, getValueFromParsedObject, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        margin: "1rem 0.5rem",
    },
    headerContainer: {
        marginBottom: "0.5rem"
    },
    header: {
        ...typography.heading1
    },
    expandedTitle: {
        color: colors.primary01Color100,
        textAlign: "right",
        display: "inline",
        marginRight: "0.5rem",
        ...typography.body1
    },
    expandedImage: {
        maxWidth: "1rem"
    },
    addonsContainer: {
        backgroundColor: colors.background02Color,
        borderRadius: "5px",
        paddingTop: "0.5rem"
    },
    clickableCard: {
        cursor: "pointer",
        width: "100%",
        margin: "0.5rem",
        padding: "1rem",
        textAlign: "center"
    }
}));

const AddonExpandable = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [childComponents, setChildComponents] = useState(null);
    const [widgetData, setWidgetData] = useState({});
    const [editing, setEditing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
    const [clickableCards, setClickableCards] = useState([]);
    // selectedCard state is used to track which AddonExpandableSection component is being selected
    const [selectedCard, setSelectedCard] = useState(null)
    const classes = useStyles();

    // If widget is selected for editing, render EditorContainer
    useEffect(() => {
        // Note that if a selectedCard is set, this AddonExpandable component will not update the editorComponent since the selected AddonExpandableSection will be the one updating it instead.
        if (shouldRenderEditorComponent && !selectedCard && widgetData.length != 0) {
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
                    clickableCards={clickableCards}
                />
            )
        }
    }, [shouldRenderEditorComponent, widgetData, saving, errorMessage, selectedCard])

    // resetSelectedCard is passed as a prop to AddonExpandableSection to unselect itself.
    const resetSelectedCard = () => {
        setSelectedCard(null)
    }

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    // It also instantiates the clickableCards state that is used to enable the user to select the card that they want to edit
    const initState = () => {
        if (!widget) return
        setChildComponents(null)
        setClickableCards([])
        setKeyValueToState(widget, "header.title", "header.title", setWidgetData)
        setKeyValueToState(widget, "header.expandedImage", "header.components[0].data.expandedImage", setWidgetData)
        setKeyValueToState(widget, "header.expandedTitle", "header.components[0].data.expandedTitle", setWidgetData)
        setChildComponents(map(widget.components, (section) => {
            // Get cardTitle for displaying in ClickableCard
            const {parsedKeyValue} = parseQUIPrefix(get(section, "data.sectionHeader.title"))
            const cardTitle = getValueFromParsedObject(parsedKeyValue)
            const clickableCard = (<Card
                key={cardTitle}
                className={clsx(classes.clickableCard)}
                onClick={()=>{
                    setSelectedCard(cardTitle)
                }}
            >
                {cardTitle}
            </Card>)
            setClickableCards(prevState => [
                ...prevState,
                clickableCard
            ])

            const shouldRenderSectionEditorComponent = selectedCard === cardTitle
            return <AddonExpandableSection 
                key={cardTitle}
                widget={section}
                updateEditorComponent={updateEditorComponent}
                shouldRenderEditorComponent={shouldRenderSectionEditorComponent}
                resetSelectedSectionCard={resetSelectedCard}
                pageId={pageId}
                fetchPageData={fetchPageData}
            />
        }))
    }

    // This effect is used to re-generate the clickableCards after a card has been clicked
    useEffect(() => {
        initState()
    }, [selectedCard])

    useEffect(() => {
        initState()
    }, [widget])

    const widgetTitle = getValueFromWidgetData(widgetData, "header.title", "header.title")
    const expandedImage = getValueFromWidgetData(widgetData, "header.expandedImage", "header.components[0].data.expandedImage")
    const expandedTitle = getValueFromWidgetData(widgetData, "header.expandedTitle", "header.components[0].data.expandedTitle")
    const headerSection = () => (<Grid container direction="row" className={clsx(classes.headerContainer)}>
        <Grid item xs={3}>
            <Typography 
                className={clsx(classes.header)} 
                variant="body1">
                    {widgetTitle}
            </Typography>
        </Grid>
        <Grid item xs={6}></Grid>
        <Grid item xs={3}>
            <Typography 
                className={clsx(classes.expandedTitle)} 
                variant="body1">
                    {expandedTitle}
            </Typography>
            <img src={expandedImage} alt="tooltip" className={clsx(classes.expandedImage)} />
        </Grid>
    </Grid>)

    return (widgetData && <Box className={clsx(classes.widgetContainer)}>
        <Grid container direction="row">
            {headerSection()}
            <Grid item xs={12} className={clsx(classes.addonsContainer)}>
                { map(childComponents, childComponent => childComponent)}
            </Grid>
        </Grid>
    </Box>)
}

export default AddonExpandable;