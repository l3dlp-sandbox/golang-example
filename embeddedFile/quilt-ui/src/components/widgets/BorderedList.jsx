import React, { useEffect, useState } from 'react';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import { get, map } from 'lodash';

import BorderedListItem from './BorderedListItem';
import { parseQUIPrefix, getValueFromWidgetData, getValueFromParsedObject, setKeyValueToState, renderEditorFieldComponents, updateLocalisationValue } from '../../helper'
import { typography, colors } from '../../styles/globalStyles'
import EditorContainer from '../EditorContainer'

const useStyles = makeStyles(() => ({
    widgetContainer: {
        margin: "1rem",
        border: `2px solid ${colors.primary01Color100}`,
        borderRadius: "0.4rem",
        boxShadow: `5px 5px ${colors.primary01Color100}`
    },
    header: {
        margin: "1rem 0px",
        ...typography.heading3,
        textAlign: "center",
        color: colors.general01Color,
        padding: "0.2rem 0px",
        backgroundImage: colors.gradient01,
        width: "100.5%",
        marginLeft: "-1px"
    },
    tileProgressTrackerContainer: {
        margin: "0.75rem",
    },
    tileImage: {
        maxWidth: "80%"
    },
    rewardImage: {
        float: "right",
        maxWidth: "1rem"
    },
    rewardText: {
        color: colors.text02Color,
        ...typography.body2,
        textAlign: "left"
    },
    tileTitle: {
        ...typography.heading3,
        marginBottom: "0.5rem"
    },
    tileProgressionText: {
        ...typography.body2,
    },
    clickableCard: {
        cursor: "pointer",
        width: "100%",
        margin: "0.5rem",
        padding: "1rem",
        textAlign: "center"
    }
}));

const BorderedList = ({widget, updateEditorComponent, resetEditorComponent, shouldRenderEditorComponent, pageId, fetchPageData}) => {
    const [childComponents, setChildComponents] = useState([]);
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
        // Note that if a selectedCard is set, this BorderedList component will not update the editorComponent since the selected BorderedListItem will be the one updating it instead.
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

    // resetSelectedCard is passed as a prop to BorderedListItem to unselect itself.
    const resetSelectedCard = () => {
        setSelectedCard(null)
    }

    // initState initializes the widgetData state for use to render the preview as well as the EditorContainer
    // It also instantiates the clickableCards state that is used to enable the user to select the card that they want to edit
    const initState = () => {
        if (!widget || get(widget, "components.length") === 0) return
        setChildComponents(null)
        setClickableCards([])
        setKeyValueToState(widget, "header", "header.title", setWidgetData)
        setChildComponents(map(widget.components, (component) => {
            // Get cardTitle for displaying in ClickableCard
            let cardTitle = ""
            const componentType = get(component, "type")
            if (componentType === "tile-progress-tracker") {
                const {parsedKeyValue} = parseQUIPrefix(get(component, "data.title"))
                cardTitle = getValueFromParsedObject(parsedKeyValue)
            } else if (componentType === "tile-disclaimer") {
                const {parsedKeyValue} = parseQUIPrefix(get(component, "data.text"))
                cardTitle = getValueFromParsedObject(parsedKeyValue)
            }

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
            const shouldRenderComponentEditorComponent = selectedCard === cardTitle
            return <BorderedListItem
                key={cardTitle}
                widget={component}
                updateEditorComponent={updateEditorComponent}
                shouldRenderComponentEditorComponent={shouldRenderComponentEditorComponent}
                resetSelectedComponentCard={resetSelectedCard}
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

    if (get(widget, "components.length") ===0) {
        return <div>Empty {widget.type}</div>
    }

    const headerTitle = getValueFromWidgetData(widgetData, "header", "header.title")
    return (widgetData && <Box className={clsx(classes.widgetContainer)}>
        <Box>
            <Typography 
                className={clsx(classes.header)} 
                variant="body1">
                    {headerTitle}
            </Typography>
        </Box>
        <Grid container direction="row">
            {map(childComponents, childComponent => childComponent)}
        </Grid>
    </Box>)
}

export default BorderedList;